package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	list := make([]string, n)

	for i := 1; i <= n; i += 1 {
		if i % 3 == 0 && i % 5 == 0 {
			list[i - 1] = "FizzBuzz"
		} else if i % 3 == 0 {
			list[i - 1] = "Fizz"
		} else if i % 5 == 0 {
			list[i - 1] = "Buzz"
		} else {
			list[i - 1] = strconv.Itoa(i)
		}
	}
	return (list)
}

func main() {
	fmt.Println(fizzBuzz(0))
	fmt.Println(fizzBuzz(1))
	// fmt.Println(fizzBuzz(-1))
	fmt.Println(fizzBuzz(15))
}
