package main

import "fmt"

func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	B := make([][]int, n)
	for i := 0; i < n; i++ {
		B[i] = make([]int, m)
	}
	fmt.Println(B)
	return B
}

