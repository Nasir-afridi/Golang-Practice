package main

import "fmt"

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("kid")
	}

	var role = "admin"
	var hasPermissions = true

	// Logical Operators
	// Or Operator
	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}
}
