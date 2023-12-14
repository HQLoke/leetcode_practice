package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	firstHalf := x
	secondHalf := 0

	if x < 0 || (x != 0 && x % 10 == 0) {
		return (false)
	} else {
		for firstHalf > secondHalf {
			secondHalf = secondHalf * 10 + firstHalf % 10
			firstHalf /= 10
		}
	}
    return (firstHalf == secondHalf || firstHalf == secondHalf / 10)
}

func main() {
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(1234))
	fmt.Println(isPalindrome(1234321))
	fmt.Println(isPalindrome(12344321))
}
