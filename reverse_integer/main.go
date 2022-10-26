package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	n := int(math.Abs(float64(x)))
	res := 0
	for n > 0 {
		res *= 10
		res += n % 10
		n /= 10
	}
	if res > ((2 << 30) - 1) {
		return 0
	}
	return int(math.Copysign(float64(res), float64(x)))
}

func main() {
	fmt.Println(reverse(-123))
}
