package main

func hammingWeight(n uint32) int {
	var m uint32 = 1
	count := 0
	for i := 0; i < 32; i++ {
		if n & m != 0 {
			count++
		}
		m = m << 1
	}
	return count
}
