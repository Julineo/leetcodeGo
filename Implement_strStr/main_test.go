package main

import "testing"

func TestStrStr(t *testing.T) {
	var tests = []struct {
		haystack string
		needle string
		want int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaa", "aaaa", -1},
		{"mississippi", "issipi", -1},
		{"a", "a", 0},
		{"","", 0}}

	for _, test := range tests {
		if got := strStr(test.haystack, test.needle); got != test.want{
			t.Errorf("strStr(%v, %v) = %v", test.haystack, test.needle, got)
		}
	}
}
