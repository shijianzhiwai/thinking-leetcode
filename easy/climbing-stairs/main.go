package main

import "fmt"

// https://leetcode-cn.com/problems/climbing-stairs/

func main() {
	fmt.Println(climbStairs2(3))
}

var ff = map[int]int{}

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}
	var ret1, ret2 int
	var ok bool
	if ret1, ok = ff[n-1]; !ok {
		ret1 = climbStairs(n - 1)
		ff[n-1] = ret1
	}
	if ret2, ok = ff[n-2]; !ok {
		ret2 = climbStairs(n - 2)
		ff[n-2] = ret2
	}
	return ret1 + ret2
}

// f(n) = f(n-1) + f(n-2)
// f(0)=1 f(1)=1 f(2)=2 f(3) 这些循环依次求出
func climbStairs2(n int) int {
	if n == 0 {
		return 1
	}
	var ret, r1, r2 int
	r2 = 1
	for i := 1; i <= n; i++ {
		ret = r1 + r2
		r1 = r2
		r2 = ret
	}
	return ret
}
