package main

import (
	"fmt"
)

type stack []rune

func (s stack) Push(r rune) (stack) {
	return append(s, r)
}

func (s stack) Top() (rune) {
	l := len(s)
	
	if l != 0 {
		return s[l - 1]
	} else {
		return rune(0)
	}
}

func (s stack) Pop() (stack) {
	l := len(s)
	return s[:l - 1]
}

func (s stack) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	st := make(stack, 0)

	matches := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i += 1 {
		val, exist := matches[rune(s[i])]
		if exist == false {
			st = st.Push(rune(s[i]))
		} else {
			if val == st.Top() {
				st = st.Pop()
			} else {
				return (false)
			}
		}
	}
	return st.IsEmpty()
}

func main() {
	fmt.Println(isValid("(")) // false
	fmt.Println(isValid("()")) // true
	fmt.Println(isValid("(]")) // false
	fmt.Println(isValid("(){}[]")) // true
	fmt.Println(isValid("()")) // true
	fmt.Println(isValid("(())())")) // false
	fmt.Println(isValid("{[()]}")) // true
	fmt.Println(isValid("(")) // false
	fmt.Println(isValid("{[()]}")) // true
	fmt.Println(isValid("(){}[]")) // true
	fmt.Println(isValid(")(")) // false
	fmt.Println(isValid("{[()]({})}")) // true
}
