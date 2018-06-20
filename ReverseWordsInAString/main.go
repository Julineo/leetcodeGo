package main

import "strings"

func reverseWords(s string) string {
	//reverse sting
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	//reverse words in string
	r := ""
	for _, v := range strings.Split(string(runes), " ") {
		runes := []rune(v)
		for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		r = r + " " + string(runes)
	}
	return strings.Trim(r, " ")
}
