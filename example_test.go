package crowd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "fmt"

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewAPI() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	api.SetUserAgent("MyApp", "1.2.3")

	user, err := api.GetUser("john", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%#v\n", user)
}

func ExampleSimplifyAttributes() {
	attrs := Attributes{
		&Attribute{Name: "test", Values: []string{"1", "2"}},
		&Attribute{Name: "abcd", Values: []string{"test", "100"}},
	}

	attrsMap := SimplifyAttributes(attrs)

	fmt.Println(attrsMap["test"])
	// Output: 1 2
}

func ExampleAPI_SetUserAgent() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	api.SetUserAgent("MyApp", "1.2.3")
}

func ExampleAPI_GetUser() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	user, err := api.GetUser("john", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%#v\n", user)
}

func ExampleAPI_Login() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	currentUser, err := api.Login("john", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%#v\n", currentUser)
}

func ExampleAPI_GetUserAttributes() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	attrs, err := api.GetUserAttributes("john")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	attrsMap := SimplifyAttributes(attrs)

	fmt.Printf("%#v\n", attrsMap)
}

func ExampleAPI_GetUserGroups() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	groups, err := api.GetUserGroups("john", GROUP_DIRECT)
	// with listing options
	groups, err = api.GetUserGroups(
		"john", GROUP_DIRECT,
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(groups) > 0 {
		for _, group := range groups {
			fmt.Printf("%#v\n", group)
		}
	}
}

func ExampleAPI_GetUserDirectGroups() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	groups, err := api.GetUserDirectGroups("john")
	// with listing options
	groups, err = api.GetUserDirectGroups(
		"john",
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(groups) > 0 {
		for _, group := range groups {
			fmt.Printf("%#v\n", group)
		}
	}
}

func ExampleAPI_GetUserNestedGroups() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	groups, err := api.GetUserNestedGroups("john")
	// with listing options
	groups, err = api.GetUserNestedGroups(
		"john",
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(groups) > 0 {
		for _, group := range groups {
			fmt.Printf("%#v\n", group)
		}
	}
}

func ExampleAPI_GetGroup() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	group, err := api.GetGroup("my_group", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%#v\n", group)
}

func ExampleAPI_GetGroupAttributes() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	attrs, err := api.GetGroupAttributes("my_group")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	attrsMap := SimplifyAttributes(attrs)

	fmt.Printf("%#v\n", attrsMap)
}

func ExampleAPI_GetGroupUsers() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	users, err := api.GetGroupUsers("my_group", GROUP_DIRECT)
	// with listing options
	users, err = api.GetGroupUsers(
		"my_group", GROUP_DIRECT,
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(users) > 0 {
		for _, user := range users {
			fmt.Printf("%#v\n", user)
		}
	}
}

func ExampleAPI_GetGroupDirectUsers() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	users, err := api.GetGroupDirectUsers("my_group")
	// with listing options
	users, err = api.GetGroupDirectUsers(
		"my_group",
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(users) > 0 {
		for _, user := range users {
			fmt.Printf("%#v\n", user)
		}
	}
}

func ExampleAPI_GetGroupNestedUsers() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	users, err := api.GetGroupNestedUsers("my_group")
	// with listing options
	users, err = api.GetGroupNestedUsers(
		"my_group",
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(users) > 0 {
		for _, user := range users {
			fmt.Printf("%#v\n", user)
		}
	}
}

func ExampleAPI_GetMemberships() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	membership, err := api.GetMemberships()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, group := range membership {
		fmt.Printf("Group: %v\n", group.Group)

		for _, userInfo := range group.Users {
			fmt.Printf(" - %v\n", userInfo.Name)
		}

		fmt.Println("")
	}
}

func ExampleAPI_SearchUsers() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	users, err := api.SearchUsers(`createdDate > 2010-12 or firstName = Jo*`)
	// with listing options
	users, err = api.SearchUsers(
		`createdDate > 2010-12 or firstName = Jo*`,
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(users) > 0 {
		for _, user := range users {
			fmt.Printf("%#v\n", user)
		}
	}
}

func ExampleAPI_SearchGroups() {
	api, err := NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	groups, err := api.SearchGroups(`name = "admin*" and active = false`)
	// with listing options
	groups, err = api.SearchGroups(
		`name = "admin*" and active = false`,
		ListingOptions{StartIndex: 100, MaxResults: 50},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(groups) > 0 {
		for _, group := range groups {
			fmt.Printf("%#v\n", group)
		}
	}
}
