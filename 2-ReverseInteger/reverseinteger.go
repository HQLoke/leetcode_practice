package main

import (
	"fmt"
)

func reverse(x int) int {
	temp := x
	sum	:= 0
	sign := 1

	if x < 0 {
		sign = -1
		temp *= -1
	}
	for temp > 0 {
		sum = sum * 10 + temp % 10
		temp /= 10
	}

	sum = sum * sign
	if x == 0 || sum > 2147483647 || sum < -2147483648 {
		return (0)
	} else {
		return (sum)
	}
}

func main() {
	fmt.Println(reverse(0))
	fmt.Println(reverse(42))
	fmt.Println(reverse(-42))
	fmt.Println(reverse(2147483647)) // 0
	fmt.Println(reverse(-2147483648)) // 0
}