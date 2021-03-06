// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	. "github.com/pingcap/check"
	"github.com/insionng/zenpress/libraries/pingcap/tidb/mysql"
	"github.com/insionng/zenpress/libraries/pingcap/tidb/util/testleak"
)

var _ = Suite(&testUtilSuite{})

type testUtilSuite struct {
}

func (s *testUtilSuite) TestDumpBinaryTime(c *C) {
	defer testleak.AfterTest(c)()
	t, err := mysql.ParseTimestamp("0000-00-00 00:00:00.0000000")
	c.Assert(err, IsNil)
	d := dumpBinaryDateTime(t, nil)
	c.Assert(d, DeepEquals, []byte{11, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0})
	t, err = mysql.ParseDatetime("0000-00-00 00:00:00.0000000")
	c.Assert(err, IsNil)
	d = dumpBinaryDateTime(t, nil)
	c.Assert(d, DeepEquals, []byte{11, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0})

	t, err = mysql.ParseDate("0000-00-00")
	c.Assert(err, IsNil)
	d = dumpBinaryDateTime(t, nil)
	c.Assert(d, DeepEquals, []byte{4, 1, 0, 1, 1})

	myDuration, err := mysql.ParseDuration("0000-00-00 00:00:00.0000000", 6)
	c.Assert(err, IsNil)
	d = dumpBinaryTime(myDuration.Duration)
	c.Assert(d, DeepEquals, []byte{0})
}
