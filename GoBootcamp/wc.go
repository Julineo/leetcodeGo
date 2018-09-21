/*
http://www.golangbootcamp.com/book/collection_types#uid98

Implement WordCount.
*/
package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
