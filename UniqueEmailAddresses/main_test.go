package main

import "testing"

func TestNumUniqueEmails(t *testing.T) {
	var tests = []struct {
		emails []string
		want int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
		}
	for _, test := range tests {
		if got := numUniqueEmails(test.emails); got != test.want {
			t.Errorf("numUniqueEmails(%v) = %v", test.emails, got)
		}
	}
}
