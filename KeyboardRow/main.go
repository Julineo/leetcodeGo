/*Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

Example 1:
Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]
Note:
You may use one character in the keyboard more than once.
You may assume the input string will only contain letters of alphabet.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	res := []string{}
	a1 := "qwertyuiop"
	a2 := "asdfghjkl"
	a3 := "zxcvbnm"
	for _, w := range words {
		count := 0
		for _, s := range strings.ToLower(w) {
			for _, a := range a1 {
				if a == s {
					count++
					break
				}
			}
		}
		if count == len(w) {
			res = append(res, w)
		}
		if count > 0 {
			continue
		}
		
		count = 0
		for _, s := range strings.ToLower(w) {
			for _, a := range a2 {
				if a == s {
					count++
					break
				}
			}
		}
		if count == len(w) {
			res = append(res, w)
		}
		if count > 0 {
			continue
		}
		
		count = 0
		for _, s := range strings.ToLower(w) {
			for _, a := range a3 {
				if a == s {
					count++
					break
				}
			}
		}
		if count == len(w) {
			res = append(res, w)
		}
	}
	return res
}
