package main

import "fmt"

// First method - convert to digit, add, then convert back to integer array
//! WRONG - Overflow issue
// func plusOne(digits []int) []int {
// 	num := 0
// 	newList := make([]int, 0)

// 	for i := 0; i < len(digits); i += 1 {
// 		num = num * 10 + digits[i]
// 	}

// 	num += 1
// 	for num > 0 {
// 		newList = append(newList, num % 10)
// 		num /= 10
// 	}
// 	for i, j := 0, len(newList)-1; i < j; i, j = i+1, j-1 {
// 		newList[i], newList[j] = newList[j], newList[i]
// 	}
// 	return newList
// }

// Second method - adding using iteration and indexing
func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	
	if digits[len(digits)-1] <= 9 {
		return digits
	}

	for i := len(digits)-1; i >= 0; i -= 1 {
		if digits[i] > 9 {
			if i != 0 {
				digits[i] -= 10
				digits[i - 1] += 1
			} else {
				digits[i] -= 10
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}

func main() {
	list := []int{1, 2, 3}
	list2 := []int{4, 5, 6}
	list3 := []int{7, 8, 9}
	list4 := []int{9, 9, 9, 9, 9, 9}

	fmt.Println(plusOne(list))
	fmt.Println(plusOne(list2))
	fmt.Println(plusOne(list3))
	fmt.Println(plusOne(list4))
}
