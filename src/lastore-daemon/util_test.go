/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package main

/*
import "testing"
import C "gopkg.in/check.v1"
import "internal/system"
import "fixme/pkg_recommend"
import "internal/system/apt"
import "strings"

type testWrap struct{}

func Test(t *testing.T) { C.TestingT(t) }
func init() {
	C.Suite(&testWrap{})
}

func (*testWrap) TestTranisition(c *C.C) {
	var data = []struct {
		from  system.Status
		to    system.Status
		valid bool
	}{
		{system.ReadyStatus, system.ReadyStatus, false},
		{system.ReadyStatus, system.RunningStatus, true},
		{system.ReadyStatus, system.FailedStatus, false},
		{system.ReadyStatus, system.SucceedStatus, false},
		{system.ReadyStatus, system.PausedStatus, true},
		{system.ReadyStatus, system.EndStatus, false},

		{system.RunningStatus, system.ReadyStatus, false},
		{system.RunningStatus, system.RunningStatus, false},
		{system.RunningStatus, system.FailedStatus, true},
		{system.RunningStatus, system.SucceedStatus, true},
		{system.RunningStatus, system.PausedStatus, true},
		{system.RunningStatus, system.EndStatus, false},

		{system.FailedStatus, system.ReadyStatus, true},
		{system.FailedStatus, system.RunningStatus, false},
		{system.FailedStatus, system.FailedStatus, false},
		{system.FailedStatus, system.SucceedStatus, false},
		{system.FailedStatus, system.PausedStatus, false},
		{system.FailedStatus, system.EndStatus, true},

		{system.SucceedStatus, system.ReadyStatus, false},
		{system.SucceedStatus, system.RunningStatus, false},
		{system.SucceedStatus, system.FailedStatus, false},
		{system.SucceedStatus, system.SucceedStatus, false},
		{system.SucceedStatus, system.PausedStatus, false},
		{system.SucceedStatus, system.EndStatus, true},

		{system.PausedStatus, system.ReadyStatus, true},
		{system.PausedStatus, system.RunningStatus, false},
		{system.PausedStatus, system.FailedStatus, false},
		{system.PausedStatus, system.SucceedStatus, false},
		{system.PausedStatus, system.PausedStatus, false},
		{system.PausedStatus, system.EndStatus, true},

		{system.EndStatus, system.ReadyStatus, false},
		{system.EndStatus, system.RunningStatus, false},
		{system.EndStatus, system.FailedStatus, false},
		{system.EndStatus, system.SucceedStatus, false},
		{system.EndStatus, system.PausedStatus, false},
		{system.EndStatus, system.EndStatus, false},
	}

	for _, d := range data {
		if !c.Check(ValidTransitionJobState(d.from, d.to), C.Equals, d.valid) {
			c.Logf("Transition %s to %s failed (%v)\n", d.from, d.to, d.valid)
		}
	}
}


func (*testWrap) TestGetEnhancedLocalePackages(c *C.C) {
	if !system.QueryPackageInstalled("deepin-desktop-base") {
		c.Skip("deepin-desktop-base not installed")
		return
	}
	lang := "tr_TR.UTF-8"

	positive := []string{"firefox-esr", "libreoffice", "thunderbird", "gimp", "chromium"}
	negative := []string{"vim"}

	for _, p := range positive {
		d := pkg_recommend.GetEnhancedLocalePackages(lang, p)
		c.Check(len(d), C.Not(C.Equals), 0)
	}
	for _, p := range negative {
		d := pkg_recommend.GetEnhancedLocalePackages(lang, p)
		c.Check(len(d), C.Equals, 0)
	}
}

func (*testWrap) TestGuestJobType(c *C.C) {
	jm := NewJobManager(apt.New(), nil)

	job, err := jm.CreateJob(system.DistUpgradeJobType, system.InstallJobType, nil)
	c.Check(err, C.Equals, nil)
	c.Check(jm.findJobByType(system.DistUpgradeJobType, nil), C.Equals, (*Job)(nil))

	job, err = jm.CreateJob("", system.DistUpgradeJobType, nil)
	c.Check(err, C.Equals, nil)
	c.Check(jm.findJobByType(system.DistUpgradeJobType, nil), C.Equals, job)
}

func (*testWrap) TestNormalizePackageNames(c *C.C) {
	s, err := NormalizePackageNames("a b c")
	c.Check(err, C.Equals, nil)
	c.Check(strings.Join(s, "_"), C.Equals, "a_b_c")

	s, err = NormalizePackageNames("")
	c.Check(err, C.Not(C.Equals), nil)
	c.Check(len(s), C.Equals, 0)
}
*/
