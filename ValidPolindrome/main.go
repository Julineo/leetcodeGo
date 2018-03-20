package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main () {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("ada"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	check := []rune{}
	for _, r := range(runes) {
		if unicode.IsLetter(r) {
			check = append(check, r)
		}
	}
	fmt.Println(check)
	for i, sv := range(runes) {
		if sv != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
