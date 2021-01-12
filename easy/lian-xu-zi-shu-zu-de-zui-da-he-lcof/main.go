package main

import "fmt"

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

var nums = []int{-2,1,-3,4,-1,2,1,-5,4}

// 动态规划接替思路
// 循环开始加和，生成 dp 数组，如果 dp[i-1] 为负数，那么 dp[i] = num[i]

func main() {
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	var m int
	for i := range nums {
		if i == 0 {
			m = nums[i]
			continue
		} else {
			if nums[i-1] <= 0 {
				m = max(m, nums[i])
				continue
			} else {
				nums[i] += nums[i-1]
				m = max(m, nums[i])
			}
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
