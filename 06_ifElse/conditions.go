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
	var hasPermissions = false

	// Logical Operators
	// Or Operator
	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}

	// And Operator
	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	}

	// Declaring variables direct in if condition.
	if ages := 10; ages >= 18 {
		fmt.Println("Person is an adult")
	} else if ages >= 12 {
		fmt.Println("Person is an Teenager")
	} else {
		fmt.Println("Person is kid")
	}

	// Go does not have ternary operator, we will have to use normal if else.
}
