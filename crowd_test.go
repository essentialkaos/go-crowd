package crowd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type CrowdSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CrowdSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CrowdSuite) TestBasic(c *C) {
	api, err := NewAPI("https://domain.com", "john", "qwerty1234")

	c.Assert(api, NotNil)
	c.Assert(err, IsNil)
}

func (s *CrowdSuite) TestAttributesHelpers(c *C) {
	attrs := Attributes{
		&Attribute{"test", []string{"AB", "CD"}},
		&Attribute{"magic", []string{"ABCD"}},
	}

	c.Assert(attrs.Has("unknown"), Equals, false)
	c.Assert(attrs.Has("test"), Equals, true)

	c.Assert(attrs.GetList("unknown"), HasLen, 0)
	c.Assert(attrs.GetList("test"), HasLen, 2)

	c.Assert(attrs.Get("unknown"), Equals, "")
	c.Assert(attrs.Get("test"), Equals, "AB CD")
	c.Assert(attrs.Get("magic"), Equals, "ABCD")
}
