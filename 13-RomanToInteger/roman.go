package main

import (
	"fmt"
)

func romanToInt(s string) int {
	sum := 0

	romanValues := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i += 1 {
		val := romanValues[s[i]]

		if i + 1 < len(s) && val < romanValues[s[i + 1]] {
			sum -= val
		} else {
			sum += val
		}
	}
	return (sum)
}

func main() {
	fmt.Println(romanToInt("IV"))
}
