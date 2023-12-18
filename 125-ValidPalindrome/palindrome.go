package main

import (
	"fmt"
	"strings"
)

//! beats 51% in runtime, beats 25.50% in memory
// func isPalindrome(s string) bool {
//     trimmed := make([]rune, 0)
	
// 	s = strings.ToLower(s)
// 	for i := 0; i < len(s); i += 1 {
// 		c := rune(s[i])

// 		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9' ) {
// 			trimmed = append(trimmed, c)
// 		}
// 	}
	
// 	for i, j := 0, len(trimmed) - 1; i < j; i, j = i + 1, j - 1 {
// 		if trimmed[i] != trimmed[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

//! beats 65% in runtime, beats 61% in memory
func isPalindrome(s string) bool {	
	s = strings.ToLower(s)

	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		for i < j && !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z')) {
			i += 1
		}
		for i < j && !((s[j] >= '0' && s[j] <= '9') || (s[j] >= 'a' && s[j] <= 'z')) {
			j -= 1
		}
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car")) // false
	fmt.Println(isPalindrome(" ")) // true
	fmt.Println(isPalindrome("0P")) // false
	fmt.Println(isPalindrome(".,")) // true
}
