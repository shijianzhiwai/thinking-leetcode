package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/coin-change/

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 100))
}

func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	return dp(cache, coins, amount)
}

func dp(cache map[int]int, coins []int, amount int) int {
	if ret, ok := cache[amount]; ok {
		return ret
	}

	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	minCount := math.MaxInt32
	for _, c := range coins {
		cc := dp(cache, coins, amount-c)
		if cc == -1 {
			continue
		}
		minCount = min(minCount, 1+cc)
	}
	if minCount == math.MaxInt32 {
		cache[amount] = -1
		return -1
	}
	cache[amount] = minCount
	return minCount
}

func min(n, m int) int {
	if n > m {
		return m
	}
	return n
}
