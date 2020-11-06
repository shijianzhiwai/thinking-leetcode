package main

import "fmt"

// https://leetcode-cn.com/problems/maximal-rectangle/

// matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]

var testData = [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

func main() {
	fmt.Println(maximalRectangle(testData))
}

func maximalRectangle(matrix [][]byte) int {
	var ret int

	dp := make([][]int, len(matrix))
	for i, l1 := range matrix {
		for _, l2 := range l1 {
			s := l2 - '0'
			dp[i] = append(dp[i], int(s))
		}
	}

	fmt.Println(dp)

	dpw := make([][]int, len(matrix))

	for i, l1 := range dp {
		for j, l2 := range l1 {
			if j == 0 {
				dpw[i] = append(dpw[i], l2)
				continue
			}
			if l2 == 1 {
				dpw[i] = append(dpw[i], 1+l1[j-1])
			} else {
				dpw[i] = append(dpw[i], 0)
			}
		}
	}
	fmt.Println(dpw)
	return ret
}
