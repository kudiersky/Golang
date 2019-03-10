// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

// package main

// import "fmt"

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jake".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	userName, password := args[1], args[2]

	if userName != "Jake" {
		fmt.Printf("Access denied for %q.\n", userName)
	} else if password != "1987" {
		fmt.Printf("Invalid password for %q.\n", userName)
	} else {
		fmt.Printf("Access granted to %q.\n", userName)
	}
}
