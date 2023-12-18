package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	last := len(nums) - 1

	for i := 0; i < len(nums); i += 1 {
		if nums[i] == val {
			for nums[last] == val && i < last{
				last -= 1
			}
			nums[i], nums[last] = nums[last], nums[i]
		}

		if nums[i] != val {
			k += 1
		}
	}
	return k
}

func main() {
	trial1 := []int{3,2,2,3}
	trial2 := []int{0,1,2,2,3,0,4,2}
	trial3 := []int{2,1,2,1,2,1,2,1,2}

	fmt.Println(removeElement(trial1, 3))
	fmt.Println(trial1)
	fmt.Println(removeElement(trial2, 2))
	fmt.Println(trial2)
	fmt.Println(removeElement(trial3, 2))
	fmt.Println(trial3)
}
