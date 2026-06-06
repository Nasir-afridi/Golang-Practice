package main

import "fmt"

func main() {
	// simple switch : if we have multiples conditions to check we use switch.
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}
}
