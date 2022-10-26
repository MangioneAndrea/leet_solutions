package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	n := int(math.Abs(float64(x)))
	res := 0
	for n > 0 {
		res *= 10
		res += n % 10
		n /= 10
	}
	return res == x
}

func main() {
	fmt.Println(isPalindrome(14341))
}
