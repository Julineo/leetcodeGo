package main

import "testing"

func TestToLowerCase(t *testing.T) {
	var tests = []struct {
		str string
		want string

	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
		{"H7#Il&", "h7#il&"}}
	for _, test := range tests {
		if got := toLowerCase(test.str); got != test.want {
			t.Errorf("toLowerCase(%v) = %v", test.str, got)
		}
	}
}
