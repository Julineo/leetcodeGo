package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{}, ""}}

	for _, test := range tests {
		if got := longestCommonPrefix(test.strs); got != test.want{
			t.Errorf("longestCommonPrefix(%v) = %v", test.strs, got)
		}
	}
}
