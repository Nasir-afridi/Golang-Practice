package main

import "fmt"

func main() {
	// even odd checker
	num := 17
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// grading checker
	marks := 75
	if marks >= 90 {
		fmt.Println("A")
	} else if marks >= 80 && marks < 90 {
		fmt.Println("B")
	} else if marks >= 70 && marks < 80 {
		fmt.Println("c")
	} else {
		fmt.Println("fail")
	}
}
