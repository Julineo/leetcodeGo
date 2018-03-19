package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(isAnagram("hello", "eholl"))

}

func isAnagram(s string, t string) bool {
	ms := make(map[rune]int)
	mt := make(map[rune]int)

	for _, vs := range s {
		ms[vs]++
	}

	for _, vt := range t {
		mt[vt]++
	}

	if reflect.DeepEqual(ms, mt) {
		return true
	}

	return false
}
