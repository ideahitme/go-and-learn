package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var x int

	// x & (x-1)
	// rightmost 1 to 0
	x = 5
	x &= (x - 1)
	assert(x, 4)
	// check if x is a power of two
	x = 256
	x &= (x - 1)
	assert(x, 0)

	// x | (x+1)
	// rightmost 0 to 1
	x = 5
	x |= (x + 1)
	assert(x, 7)
	// check if x = 2^n - 1
	x = 127
	assert(2*x+1, x|(x+1))

	// x & (x + 1)
	// turn all trailing 1s to 0s
	x = 11 // 0b1011
	x &= (x + 1)
	assert(x, 8)

	// x | (x-1)
	// turn all trailing 0s to 1s
	x = 12 // 0b1100
	x |= (x - 1)
	assert(x, 15)

	// bitwise negation
	fmt.Println(^2) // 00000010 -> (11111101) = -3 is a two's complement of 3,
	// because 3 = 00000011 + 11111101 = 2 ^ n
	fmt.Println(^5) // 00000101 -> (1111010) = -6
	// therefore ~x = -x - 1, for any signed integer x
	fmt.Println(^-5) // = 4

	fmt.Println(reverseBits(0))
	fmt.Println(singleNumber([]int{1, 1, 1, 3, 3, 3, 5}))
	fmt.Println(divide(-1, 1))
	fmt.Println(divide(15, 1))
	fmt.Println(divide(15, 30))
	fmt.Println(divide(45, 3))
	fmt.Println(divide(2147483648, 1))

}

// reverses bits of the integer 32bit
func reverseBits(x int) int {
	i := 32
	res := 0
	for i > 0 {
		res = res << 1
		res |= x & 1
		x = x >> 1
		i--
	}
	return res
}

// find an integer in the array which occurs only once, if all other integers occur three times
func singleNumber(A []int) int {
	bits := make([]int, 32)
	for i := 0; i < 32; i++ {
		total := 0
		for j := 0; j < len(A); j++ {
			total += A[j] & 1
			A[j] = A[j] >> 1
		}
		bits[i] = total % 3
	}
	result := 0
	x := 1
	for i := 0; i < 32; i++ {
		result += x * bits[i]
		x = x << 1
	}
	return result
}

func assert(x, y int) {
	if x != y {
		log.Fatalf("%d != %d", x, y)
	}
}

func countOnes(x int) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

// divide one number by the other without using multiplication, modulus, or division operator
func divide(x, y int) int {
	if y == 0 {
		return math.MaxInt32
	}
	signed := false
	if x < 0 && y > 0 {
		x = -x
		signed = true
	}
	if x > 0 && y < 0 {
		y = -y
		signed = true
	}
	if x < 0 && y < 0 {
		x = -x
		y = -y
	}
	answer := 0
	calc := 1
	for x > 0 {
		for y > x {
			y = y >> 1
			calc = calc >> 1
		}
		x -= y
		answer += calc
		y = y << 1
		calc = calc << 1
	}
	if answer > math.MaxInt32 && !signed {
		answer = math.MaxInt32
	}
	if signed {
		return -answer
	}
	return answer
}
