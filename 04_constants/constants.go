package main

import "fmt"

// constant and variable also working outside the main function.
const age = 30

func main() {

	// in constants we do not change the value of the variable the value is fix.
	const name = "golang"
	fmt.Println(name, age)

	//we also grouped multiple constants.
	const (
		text    = "hello"
		isAdult = true
	)
	fmt.Println(text, isAdult)

}
