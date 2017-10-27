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
	C "gopkg.in/check.v1"
	"os"
	"testing"
)

type testWrap struct{}

func Test(t *testing.T) { C.TestingT(t) }
func init() {
	C.Suite(&testWrap{})
}

func (*testWrap) TestDetect(c *C.C) {
	if os.Getenv("NO_TEST_NETWORK") == "1" {
		c.Skip("NO_TEST_NETWORK")
	}
	data := []struct {
		Official   string
		Mirror     string
		IsOfficial bool
	}{
		{
			"http://localhost/abc",
			"http://depo.pardus.org.tr/deepin/dists/onyedi/Release",
			false,
		},
	}

	for _, item := range data {
		r := MakeChoice(item.Official, item.Mirror)

		var expect string
		if item.IsOfficial {
			expect = item.Official
		} else {
			expect = item.Mirror
		}

		if !c.Check(r, C.Equals, expect) {
			c.Fatalf("Failed when %q,%q\n", item.Official, item.Mirror)
		}
	}
}
