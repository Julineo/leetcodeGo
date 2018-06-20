package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	var tests = []struct {
		s string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"", ""},
	}
	for _, test := range tests {
		if got := reverseWords(test.s); got != test.want {
			t.Errorf("reversWords(%v) = %v", test.s, got)
		}
	}

}
