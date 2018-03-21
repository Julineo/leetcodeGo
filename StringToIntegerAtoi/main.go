package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("-345 54"))
}

func myAtoi(str string) int {
	runes := []rune(str)
	digits := []rune{}
	for _, r := range(runes) {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	fmt.Println(digits)
	}
return 0
}
