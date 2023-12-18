package main

import (
	"fmt"
	// "math"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		 return x
	}

	start, end, mid := 0, x, -1

	for start <= end {
		mid = start + (end - start) / 2;

		if (mid * mid == x) {
			return mid
		} else if (mid * mid < x) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if mid * mid > x {
		return mid - 1
	}
	return mid
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(42))
	fmt.Println(mySqrt(120))
	fmt.Println(mySqrt(101010))
	fmt.Println(mySqrt(2147483647))
}
