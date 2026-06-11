package main

import "fmt"

func main() {
	// slice -> dynamic arrays means not fixed size like arrays.

	var nums []int // declared slice but it is uninitialized slice and it is nil.
	fmt.Println(nums == nil)
	fmt.Println(len(nums)) // checked the length of the slice

	// make([]int, length, capacity) creates a slice with pre-allocated memory
	// length   = number of elements accessible right now (initialized with zero values)
	// capacity = total memory reserved internally (can grow without re-allocation)

	var num = make([]int, 2, 5)
	// num = [0, 0]  <-- 2 elements accessible, but memory for 5 is already reserved

	fmt.Println(len(num)) // 2  --> only 2 elements are currently accessible
	fmt.Println(cap(num)) // 5  --> but memory for 5 elements is pre-allocated
}
