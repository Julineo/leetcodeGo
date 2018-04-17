package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxProfit([]int{1,2}))//1
	fmt.Printf("%v\n", maxProfit([]int{2,1}))//0
}

func maxProfit(prices []int) int {
    i := 0
    valley, peak, profit := 0, 0, 0
    for i < len(prices) - 1 {
        for i < len(prices) - 1 && prices[i] >= prices[i + 1] {
            peak = prices[i]
            i++
        } 
        for i < len(prices) - 1 && prices[i] <= prices[i + 1] {
            valley = prices[i]
            peak = prices[i+1]
            i++
        }
        profit = profit + peak - valley
    }
    return profit
}
