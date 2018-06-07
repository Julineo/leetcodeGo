/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/
package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pref := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		for strings.Index(str, pref) != 0 {
			pref = pref[:len(pref)-1]
			if pref == "" {
				return pref
			}
		}
	}
	return pref
}
