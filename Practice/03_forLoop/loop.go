package main

import "fmt"

func main() {
	// Print 1 to 10 using for loop
	a := 1
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	// Print 5 table
	for i := 1; i <= 10; i++ {
		fmt.Println(5, "x", i, "=", 5*i)
	}
}
