package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	answer := make([]int, 0)

	for i := 0; i < len(nums) - 1; i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] + nums[j] == target {
				answer = append(answer, i)
				answer = append(answer, j)
			}
		}
	}
	return (answer)
}

func main() {
	sample := []int{1, 2, 3}

	fmt.Println(twoSum(sample, 6))
}