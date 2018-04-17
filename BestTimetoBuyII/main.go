package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxProfit([]int{})) //0
	fmt.Printf("%v\n", maxProfit([]int{1, 2})) //1
	fmt.Printf("%v\n", maxProfit([]int{2, 1})) //0
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	i := 0
	valley, peak, profit := prices[0], prices[0], 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		profit = profit + peak - valley
	}
	return profit
}
