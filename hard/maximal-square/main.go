package main

import "fmt"

// https://leetcode-cn.com/problems/maximal-square/
// 求 DP 值

var testData = [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

func main() {
	fmt.Println(maximalRectangle(testData))
}

func maximalRectangle(matrix [][]byte) int {
	maxSize := 0
	dp := make([][]int, len(matrix))
	for k, v := range matrix {
		for _, b := range v {
			c := int(b - '0')
			if c == 1 {
				maxSize = 1
			}
			dp[k] = append(dp[k], c)
		}
	}

	for i, v := range dp {
		for j, a := range v {
			if a == 1 && i > 0 && j > 0 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSize {
					maxSize = dp[i][j]
				}
			}
		}
	}
	return maxSize * maxSize
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
