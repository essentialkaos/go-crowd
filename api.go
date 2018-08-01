package crowd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
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

// Membership contains membership info
type Membership struct {
	Group string      `xml:"group,attr"`
	Users []*UserInfo `xml:"users>user"`
}

// UserInfo contains basic user info (username)
type UserInfo struct {
	Username string `xml:"name,attr"`
}

// User contains info about user
type User struct {
	Username    string       `xml:"name,attr"`
	FirstName   string       `xml:"first-name"`
	LastName    string       `xml:"last-name"`
	DisplayName string       `xml:"display-name"`
	Email       string       `xml:"email"`
	Key         string       `xml:"key"`
	IsActive    bool         `xml:"active"`
	Attributes  []*Attribute `xml:"attributes>attribute"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String convert user info to string
func (u *UserInfo) String() string {
	return u.Username
}

// String convert attribute to string
func (a *Attribute) String() string {
	return fmt.Sprintf("%s:%v", a.Name, a.Values)
}

// Error convert crowd errro to error struct
func (e crowdError) Error() error {
	return errors.New(e.Message)
}
