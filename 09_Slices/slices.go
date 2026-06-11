package main

import "fmt"

func main() {
	// slice -> dynamic arrays means not fixed size like arrays.

	var nums []int // declared slice but it is uninitialized slice and it is nil.
	fmt.Println(nums == nil)
	fmt.Println(len(nums)) // checked the length of the slice

	// slice with make function.
	var num = make([]int, 2) // capacity is 2 means maximum numbers of elements can fit.
	fmt.Println(cap(num))
}
