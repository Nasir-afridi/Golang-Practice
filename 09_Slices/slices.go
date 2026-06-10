package main

import "fmt"

func main() {
	// slice -> dynamic arrays means not fixed size like arrays.
	var nums []int // declared slice but it is uninitialized slice and it is nil.

	fmt.Println(nums == nil)
}
