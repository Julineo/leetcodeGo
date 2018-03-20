package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main () {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0p"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	check := make([]rune,0)
	for _, r := range(runes) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			check = append(check, r)
		}
	}
	for i, sv := range(check) {
		if sv != check[len(check)-i-1] {
			return false
		}
	}
	return true
}
