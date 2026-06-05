package main

import "fmt"

func main() {
	var name string = "Nasir"
	fmt.Println(name)

	// type infer means golang itself give type to variable according to the value of the variable.
	var age = 32
	fmt.Println(age)

	//shorthand syntax: cannot working outside the function always use it inside function
	city := "islamabad"
	fmt.Println(city)

}
