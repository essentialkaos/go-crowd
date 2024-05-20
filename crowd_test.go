package crowd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
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

func (s *CrowdSuite) TestListingOptionsEncoder(c *C) {
	l1 := ListingOptions{}
	l2 := ListingOptions{MaxResults: 3}
	l3 := ListingOptions{StartIndex: 5}
	l4 := ListingOptions{StartIndex: 5, MaxResults: 7}
	l5 := ListingOptions{StartIndex: -1, MaxResults: 0}

	c.Assert(l1.Encode(), Equals, "")
	c.Assert(l2.Encode(), Equals, "&max-results=3")
	c.Assert(l3.Encode(), Equals, "&start-index=5")
	c.Assert(l4.Encode(), Equals, "&start-index=5&max-results=7")
	c.Assert(l5.Encode(), Equals, "")
}
