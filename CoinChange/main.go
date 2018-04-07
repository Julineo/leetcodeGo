package main

func coinChange(coins []int, amount int) int {
	max := amount + 1
	fs := make([]int, amount + 1)
	for i, _ := range fs {
		fs[i] = max
	}
	fs[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				fs[i] = min(fs[i], fs[i - coins[j]] +1)
			}
		}
	}

	if fs[amount] > amount {
		return -1
	}
	return fs[amount] 
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
