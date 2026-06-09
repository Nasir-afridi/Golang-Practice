package main

import "fmt"

func main() {
	// Array length will be defined by us.
	// after the nums we defined how many elements are in it.

	var nums [4]int        // declared the array.
	fmt.Println(len(nums)) // get the length of the array

	// adding the element on the 0th index.
	nums[0] = 1
	fmt.Println(nums[0])

	// declare array in single line.
	number := [5]int{1, 2, 3, 4, 5}
	fmt.Println(number)
}
