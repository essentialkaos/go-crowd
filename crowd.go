package crowd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"time"

	"github.com/erikdubbelboer/fasthttp"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// API is Confluence API struct
type API struct {
	Client *fasthttp.Client // Client is client for http requests

	url       string // confluence URL
	basicAuth string // basic auth
}

// ////////////////////////////////////////////////////////////////////////////////// //

// API errors
var (
	ErrInitEmptyURL      = errors.New("URL can't be empty")
	ErrInitEmptyUser     = errors.New("User can't be empty")
	ErrInitEmptyPassword = errors.New("Password can't be empty")
	ErrNoPerms           = errors.New("User does not have permission to use crowd")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewAPI create new API struct
func NewAPI(url, username, password string) (*API, error) {
	switch {
	case url == "":
		return nil, ErrInitEmptyURL
	case username == "":
		return nil, ErrInitEmptyUser
	case password == "":
		return nil, ErrInitEmptyPassword
	}

	return &API{
		Client: &fasthttp.Client{
			Name:                getUserAgent("", ""),
			MaxIdleConnDuration: 5 * time.Second,
			ReadTimeout:         3 * time.Second,
			WriteTimeout:        3 * time.Second,
			MaxConnsPerHost:     150,
		},

		url:       url,
		basicAuth: genBasicAuthHeader(username, password),
	}, nil
}

// SetUserAgent set user-agent string based on app name and version
func (api *API) SetUserAgent(app, version string) {
	api.Client.Name = getUserAgent(app, version)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUser returns a user
func (api *API) GetUser(userName string, withAttributes bool) (*User, error) {
	url := "rest/usermanagement/1/user?username=" + esc(userName)

	if withAttributes {
		url += "&expand=attributes"
	}

	result := &User{}
	statusCode, err := api.doRequest("GET", url, result, nil)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// GetUserAttributes returns a list of user attributes
func (api *API) GetUserAttributes(userName string) ([]*Attribute, error) {
	return api.getAttributes("rest/usermanagement/1/user/attribute?username=" + esc(userName))
}

// GetUserGroups returns the groups that the user is a member of
func (api *API) GetUserGroups(userName, groupType string) ([]*Group, error) {
	result := &struct {
		Groups []*Group `xml:"group"`
	}{}
	statusCode, err := api.doRequest(
		"GET", "rest/usermanagement/1/user/group/"+esc(groupType)+"?expand=group&username="+esc(userName),
		result, nil,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result.Groups, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// GetUserDirectGroups returns the groups that the user is a direct member of
func (api *API) GetUserDirectGroups(userName string) ([]*Group, error) {
	return api.GetUserGroups(userName, GROUP_DIRECT)
}

// GetUserNestedGroups returns the groups that the user is a nested member of
func (api *API) GetUserNestedGroups(userName string) ([]*Group, error) {
	return api.GetUserGroups(userName, GROUP_NESTED)
}

// GetGroup returns a group
func (api *API) GetGroup(groupName string, withAttributes bool) (*Group, error) {
	url := "rest/usermanagement/1/group?groupname=" + esc(groupName)

	if withAttributes {
		url += "&expand=attributes"
	}

	result := &Group{}
	statusCode, err := api.doRequest("GET", url, result, nil)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// GetGroupAttributes returns a list of group attributes
func (api *API) GetGroupAttributes(groupName string) ([]*Attribute, error) {
	return api.getAttributes("rest/usermanagement/1/group/attribute?groupname=" + esc(groupName))
}

// GetGroupUsers returns the users that are members of the specified group
func (api *API) GetGroupUsers(groupName, groupType string) ([]*User, error) {
	result := &struct {
		Users []*User `xml:"user"`
	}{}
	statusCode, err := api.doRequest(
		"GET", "rest/usermanagement/1/group/user/"+esc(groupType)+"?expand=user&groupname="+esc(groupName),
		result, nil,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result.Users, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// GetGroupDirectUsers returns the users that are direct members of the specified group
func (api *API) GetGroupDirectUsers(groupName string) ([]*User, error) {
	return api.GetGroupUsers(groupName, GROUP_DIRECT)
}

// GetGroupNestedUsers returns the users that are nested members of the specified group
func (api *API) GetGroupNestedUsers(groupName string) ([]*User, error) {
	return api.GetGroupUsers(groupName, GROUP_NESTED)
}

// GetMemberships returns full details of all group memberships, with users and
// nested groups
func (api *API) GetMemberships() ([]*Membership, error) {
	result := &struct {
		Memberships []*Membership `xml:"membership"`
	}{}
	statusCode, err := api.doRequest(
		"GET", "rest/usermanagement/1/group/membership",
		result, nil,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result.Memberships, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// SearchUsers searches for users with the specified search restriction
func (api *API) SearchUsers(cql string) ([]*User, error) {
	result := &struct {
		Users []*User `xml:"user"`
	}{}
	statusCode, err := api.doRequest(
		"GET", "rest/usermanagement/1/search?entity-type=user&&expand=user&restriction="+esc(cql),
		result, nil,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result.Users, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// SearchGroups searches for groups with the specified search restriction
func (api *API) SearchGroups(cql string) ([]*Group, error) {
	result := &struct {
		Groups []*Group `xml:"group"`
	}{}
	statusCode, err := api.doRequest(
		"GET", "rest/usermanagement/1/search?entity-type=group&&expand=group&restriction="+esc(cql),
		result, nil,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result.Groups, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getAttributes fetch attributes
func (api *API) getAttributes(url string) ([]*Attribute, error) {
	result := &struct {
		Attributes []*Attribute `xml:"attribute"`
	}{}

	statusCode, err := api.doRequest("GET", url, result, nil)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result.Attributes, nil
	case 403:
		return nil, ErrNoPerms
	default:
		return nil, makeUnknownError(statusCode)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[ARITY]

// doRequest create and execute request
func (api *API) doRequest(method, uri string, result, body interface{}) (int, error) {
	req := api.acquireRequest(method, uri)
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if body != nil {
		bodyData, err := xml.Marshal(body)

		if err != nil {
			return -1, err
		}

		req.SetBody(bodyData)
	}

	err := api.Client.Do(req, resp)

	if err != nil {
		return -1, err
	}

	statusCode := resp.StatusCode()

	if statusCode != 200 && statusCode < 500 {
		e := &crowdError{}
		err = xml.Unmarshal(resp.Body(), e)

		if err != nil {
			return statusCode, err
		}

		return statusCode, e.Error()
	}

	if result == nil {
		return statusCode, nil
	}

	err = xml.Unmarshal(resp.Body(), result)

	return statusCode, err
}

// codebeat:enable[ARITY]

// acquireRequest acquire new request with given params
func (api *API) acquireRequest(method, uri string) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(api.url + uri)

	if method != "GET" {
		req.Header.SetMethod(method)
	}

	// Set auth header
	req.Header.Add("Authorization", "Basic "+api.basicAuth)

	return req
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getUserAgent generate user-agent string for client
func getUserAgent(app, version string) string {
	if app != "" && version != "" {
		return fmt.Sprintf(
			"%s/%s %s/%s (go; %s; %s-%s)",
			app, version, NAME, VERSION, runtime.Version(),
			runtime.GOARCH, runtime.GOOS,
		)
	}

	return fmt.Sprintf(
		"%s/%s (go; %s; %s-%s)",
		NAME, VERSION, runtime.Version(),
		runtime.GOARCH, runtime.GOOS,
	)
}

// genBasicAuthHeader generate basic auth header
func genBasicAuthHeader(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

// makeUnknownError create error struct for unknown error
func makeUnknownError(statusCode int) error {
	return fmt.Errorf("Unknown error occurred (status code %d)", statusCode)
}

// esc escapes the string so it can be safely placed inside a URL query
func esc(s string) string {
	return url.QueryEscape(s)
}
