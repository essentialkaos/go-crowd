package crowd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Groups types
const (
	GROUP_DIRECT = "direct"
	GROUP_NESTED = "nested"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// crowdError is crowd error struct
type crowdError struct {
	Message string `xml:"message"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Attributes it is slice with attributes
type Attributes []*Attribute

// Attribute contains attribute info
type Attribute struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:"values>value"`
}

// Group contains info about group
type Group struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Type        string `xml:"type"`
	IsActive    bool   `xml:"active"`
}

// GroupAttributes contains group attributes
type GroupAttributes struct {
	XMLName    xml.Name     `xml:"attributes"`
	Attributes []*Attribute `xml:"attribute"`
}

// Membership contains membership info
type Membership struct {
	Group string      `xml:"group,attr"`
	Users []*UserInfo `xml:"users>user"`
}

// UserInfo contains basic user info (username)
type UserInfo struct {
	Name string `xml:"name,attr"`
}

// User contains info about user
type User struct {
	Attributes  Attributes `xml:"attributes>attribute"`
	Name        string     `xml:"name,attr"`
	FirstName   string     `xml:"first-name"`
	LastName    string     `xml:"last-name"`
	DisplayName string     `xml:"display-name"`
	Email       string     `xml:"email"`
	Key         string     `xml:"key,omitempty"`
	Password    string     `xml:"password>value,omitempty"`
	IsActive    bool       `xml:"active"`
}

// UserAttributes contains user attributes
type UserAttributes struct {
	XMLName    xml.Name     `xml:"attributes,omitempty"`
	Attributes []*Attribute `xml:"attribute"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String convert user info to string
func (u *UserInfo) String() string {
	return u.Name
}

// String convert attribute to string
func (a *Attribute) String() string {
	return fmt.Sprintf("%s:%v", a.Name, a.Values)
}

// Has returns true if slice contains attribute with given name
func (a Attributes) Has(name string) bool {
	if len(a) == 0 {
		return false
	}

	for _, attr := range a {
		if attr.Name == name {
			return true
		}
	}

	return false
}

// GetList returns slice with values for attribute with given name
func (a Attributes) GetList(name string) []string {
	if len(a) == 0 {
		return nil
	}

	for _, attr := range a {
		if attr.Name == name {
			return attr.Values
		}
	}

	return nil
}

// Get returns merged values for attribute with given name
func (a Attributes) Get(name string) string {
	values := a.GetList(name)

	if len(values) != 0 {
		return strings.Join(values, " ")
	}

	return ""
}

// Error convert crowd errro to error struct
func (e crowdError) Error() error {
	return errors.New(e.Message)
}
