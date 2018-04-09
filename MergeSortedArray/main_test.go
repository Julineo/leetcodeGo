package main

import (
	"testing"
	"reflect"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		nums1 []int
		m int
		nums2 []int
		n int
		want []int
	}{
		{
			[]int{1,2,3,0,0,0},
			3,
			[]int{2,5,6},
			3,
			[]int{1,2,2,3,5,6}},
	}
	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.want,test.nums1) {
			t.Errorf("merge want: %v, got: %v", test.want, test.nums1)
		}
	}
}

