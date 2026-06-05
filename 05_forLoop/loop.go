package main

import "fmt"

func main() {
	// For -> only construct in golang for looping. we can create while loop with the help of for loop.

	// while loop
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	for {
		fmt.Println("1")
	}

	// Classic For loop
	for a := 0; a <= 3; a++ {

		// it will break
		if a == 3 {
			break
		}

		// it will continue
		if a == 2 {
			continue
		}

		fmt.Println(a)
	}

	// Loop Using Range: it will print from 0 to 2.
	for b := range 3 {
		fmt.Println(b)
	}
}
