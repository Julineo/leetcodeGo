package main

func peakIndexInMountainArray(A []int) int {
	idx := 0
	for {
		if idx + 1 == len(A) {
			return -1
		}
		if A[idx] > A[idx+1] {
			return idx
		}
		idx++
	}
	return -1
}
