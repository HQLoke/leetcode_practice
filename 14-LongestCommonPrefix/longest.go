package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
    long := ""

	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	for i, j := 0, 0; i < len(first) && j < len(last); i, j = i + 1, j + 1 {
		if first[i] == last[j] {
			long += string(first[i])
		} else {
			break
		}
	}
	return (long)
}

func main() {
	list := []string{"apple", "approve", "antsy", "amalgamation", "alkaline"}
	list2 := []string{"dog","racecar","car"}

	fmt.Println(longestCommonPrefix(list))
	fmt.Println(longestCommonPrefix(list2))
}
