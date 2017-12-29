package main

import "fmt"

func main() {
	fmt.Println(nextPowerOfTwo(7))
	fmt.Println(nextPowerOfTwo(8))
	fmt.Println(nextPowerOfTwo(16))
	fmt.Println(nextPowerOfTwo(15))
	fmt.Println(nextPowerOfTwo(124))
}

func nextPowerOfTwo(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}
