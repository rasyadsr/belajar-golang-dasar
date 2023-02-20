package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You're Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

/*

Manual, tanpa anonymous function

func blackListAdmin(name string) bool {
	return name == "admin"
}

func blackListRoot(name string) bool {
	return name == "root"
}
*/

func main() {
	/* lebih baik seperti ini, karena cukup deklarasikan sekali function nya*/
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("messi", blackList)

	/* atau bisa langsung seperti ini*/
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}
