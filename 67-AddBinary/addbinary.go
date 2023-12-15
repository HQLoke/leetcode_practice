package main

import (
	"fmt"
)

func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 

func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	len_a := len(a) - 1
	len_b := len(b) - 1

	for len_a >= 0 || len_b >= 0 || carry > 0 {
		if len_a >= 0 {
			carry += int(a[len_a]) - '0'
			len_a -= 1
		}

		if len_b >= 0 {
			carry += int(b[len_b]) - '0'
			len_b -= 1
		}
		ans += string(carry % 2 + '0')
		carry /= 2
	}
	return reverse(ans)	
}

func main() {
	fmt.Println(addBinary("11", "1")) // "100"
	fmt.Println(addBinary("1010", "1011")) // "10101"
	fmt.Println(addBinary("0", "0")) // "0"
	fmt.Println(addBinary("1101", "100")) // "10001"
	fmt.Println(addBinary("1111", "1111")) // "11110"
	fmt.Println(addBinary("1", "111")) // "1000"
	fmt.Println(addBinary("0", "1101")) // "1101"
	fmt.Println(addBinary("101010", "1101")) // "110111"
	fmt.Println(addBinary("111000", "10101")) // "1001001"
	fmt.Println(addBinary("111111", "1")) // "1000000"
}
