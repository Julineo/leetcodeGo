/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(addBinary("11", "1"), "100")
	fmt.Println(addBinary("11", "11"), "110")
	fmt.Println(addBinary("11", "10"), "101")
	fmt.Println(addBinary("10", "10"), "100")

//	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", 
//	"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}

func addBinary(a string, b string) (res string) {
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		//48 - 0; 49 - 1
		if a[i] == 48 && b[i] == 48 {
			res = strconv.Itoa(carry + 0) + res
			carry = 0
			continue
		}
		if ((a[i] == 48 && b[i] == 49) || (a[i] == 49 && b[i] == 48)) && carry == 0 {
			res = "1" + res
			continue

		}
		if ((a[i] == 48 && b[i] == 49) || (a[i] == 49 && b[i] == 48)) && carry == 1 {
			res = "0" + res
			continue

		}

		if a[i] == 49 && b[i] == 49 {
			res =  strconv.Itoa(carry + 0) + res
			carry = 1
			continue
		}
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}
