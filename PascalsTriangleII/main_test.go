package main

import (
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {
	var tests = []struct {
		row  int
		want []int
	}{
		{3, []int{1, 3, 3, 1}},
		{2, []int{1, 2, 1}},
		{0, []int{1}},
	}
	for _, test := range tests {
		if got := getRow(test.row); !reflect.DeepEqual(got, test.want) {
			t.Errorf("getRow(%v) = %v", test.row, got)
		}
	}
}
