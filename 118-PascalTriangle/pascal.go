package main

import (
	"fmt"
	// "gonum.org/v1/gonum/stat/combin"
)

func binomial(n, k int) int {
	if k > n / 2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func generate(numRows int) [][]int {
	arrays := make([][]int, numRows)

	for i := 0; i < numRows; i += 1 {
		arrays[i] = make([]int, i + 1)
		arrays[i][0] = 1
		arrays[i][i] = 1

		for j := 1; j <= i; j += 1 {
			combi := binomial(i, j)
			arrays[i][j] = combi
		}
	}
	return arrays
}

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(5))
	fmt.Println(generate(10))
	fmt.Println(generate(22))
	fmt.Println(generate(30))
}