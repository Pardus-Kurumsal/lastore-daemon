/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"internal/system"
	"pkg.deepin.io/lib/dbus"
	"time"
)

type Job struct {
	next   *Job
	option map[string]string

	Id         string
	Name       string
	Packages   []string
	CreateTime int64

	Type string

	Status system.Status

	Progress    float64
	Description string

	// completed bytes per second
	Speed      int64
	speedMeter SpeedMeter

	Cancelable bool

	queueName string
	retry     int

	// adjust the progress range, used by some download job type
	progressRangeBegin float64
	progressRangeEnd   float64
}

func NewJob(id string, jobName string, packages []string, jobType string, queueName string) *Job {
	j := &Job{
		Id:         id,
		Name:       jobName,
		CreateTime: time.Now().UnixNano(),
		Type:       jobType,
		Packages:   packages,
		Status:     system.ReadyStatus,
		Progress:   .0,
		Cancelable: true,

		option:    make(map[string]string),
		queueName: queueName,
		retry:     3,

		progressRangeBegin: 0,
		progressRangeEnd:   1,
	}
	if jobType == system.DownloadJobType {
		go j.initDownloadSize()
	}
	return j
}

func (j *Job) initDownloadSize() {
	s, err := system.QueryPackageDownloadSize(j.Packages...)
	if err != nil {
		log.Warnf("initDownloadSize failed: %v", err)
		return
	}
	j.speedMeter.SetDownloadSize(int64(s))
}

func (j *Job) changeType(jobType string) {
	j.Type = jobType
}

func (j Job) String() string {
	return fmt.Sprintf("Job{Id:%q:%q,Type:%q(%v,%v), %q(%.2f)}@%q",
		j.Id, j.Packages,
		j.Type, j.Cancelable, j.Status,
		j.Description, j.Progress, j.queueName,
	)
}

// _UpdateInfo update Job information from info and return
// whether the information changed.
func (j *Job) _UpdateInfo(info system.JobProgressInfo) bool {
	var changed = false

	if info.Description != j.Description {
		changed = true
		j.Description = info.Description
		dbus.NotifyChange(j, "Description")
	}
	if info.Cancelable != j.Cancelable {
		changed = true
		j.Cancelable = info.Cancelable
		dbus.NotifyChange(j, "Cancelable")
	}
	log.Tracef("updateInfo %v <- %v\n", j, info)

	cProgress := buildProgress(info.Progress, j.progressRangeBegin, j.progressRangeEnd)
	if cProgress > j.Progress {
		changed = true
		j.Progress = cProgress
		dbus.NotifyChange(j, "Progress")
	}

	// see the apt.go, we scale download progress value range in [0,0.5
	speed := j.speedMeter.Speed(info.Progress)

	if speed != j.Speed {
		changed = true
		j.Speed = speed
		dbus.NotifyChange(j, "Speed")
	}

	if info.Status != j.Status {
		err := TransitionJobState(j, info.Status)
		if err != nil {
			log.Warnf("_UpdateInfo: %v\n", err)
			return false
		}
		changed = true
	}
	return changed
}

func (j *Job) _InitProgressRange(begin, end float64) {
	if end <= begin || end-begin > 1 || begin > 1 || end > 1 {
		panic("Invalid Progress range init")
	}
	if j.Progress != 0 {
		panic("InitProgressRange can only invoke once and before job start")
	}
	j.progressRangeBegin = begin
	j.progressRangeEnd = end
	j.Progress = j.progressRangeBegin
}

func buildProgress(p, begin, end float64) float64 {
	return begin + p*(end-begin)
}
