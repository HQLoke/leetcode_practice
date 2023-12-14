package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	lists := strings.Split(s, " ")
	lastWord := lists[len(lists)-1]

	return len(lastWord)
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World")) // 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy")) // 6
}
