package main

import (
	"fmt"
	"unicode"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("-9223372036854775809"))
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	minusOne := int64(1)
	if rune(str[0]) == '-' { //check if negative
		minusOne = -1
	}
	runes := []rune(str)
	digits := []int64{}
	if runes[0] == '+' || runes[0] == '-' {
		runes = runes[1:]
	}
	for _, r := range(runes) {
		if unicode.IsDigit(r) {
			digits = append(digits, int64(r)-48)
			continue
		} else {
			break
		}
		if r == ' ' && len(digits) > 0 {
			break 
		}
		if !unicode.IsDigit(r) && len(digits) > 0 {
			break
		}
	}
	if len(digits) == 0 {
		return 0
	}
	result := digits[0]
	for i := 1; i < len(digits); i++ {
		result = result * 10 + digits[i]
		if result > math.MaxInt32 + 1 {
			break
		}
	}
	result = result * minusOne
	if result < math.MinInt32 && result < 0 {
		return math.MinInt32
	}
	if result > math.MaxInt32 && result > 0 {
		return math.MaxInt32
	}
	return int(result)
}
