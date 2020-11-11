package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/

// 状态自动机 || 有限状态机

/*
			0      1    2   3
		backspace|flag|num|non-num
start => [start num num end]
flag => [end end num end]
num => [end end num end]
*/

var state [][]int

func init() {
	state = make([][]int, 3)
	// 上一个 => 如果下一个是索引值的 flag，则...
	state[start] = []int{start, num, num, end}
	state[flag] = []int{end, end, num, end}
	state[num] = []int{end, end, num, end}
}

const (
	end    = -1
	start  = 0
	flag   = 1
	num    = 2
	nonNum = 3
)

func main() {
	fmt.Println(myAtoi("  55555kalsjdkjlasjkljkdsal"))
}

func myAtoi(str string) int {
	retNum := 0
	numFlag := 1
	lastFlag := start
	for _, s := range str {
		fl := getFlag(byte(s))
		ret := state[lastFlag]
		resp := ret[fl]
		if resp == end {
			return retNum * numFlag
		}
		if fl == flag {
			if s == '-' {
				numFlag = -1
			}
		}
		if fl == num {
			retNum = retNum*10 + int(s-'0')
			if numFlag == 1 && retNum > math.MaxInt32{
				return math.MaxInt32
			}
			if numFlag == -1 && -retNum < math.MinInt32 {
				return math.MinInt32
			}
		}
		lastFlag = fl
	}
	return retNum * numFlag
}

func getFlag(bt byte) int {
	if bt == '-' || bt == '+' {
		return flag
	}
	if bt == ' ' {
		return start
	}
	if '0' <= bt && bt <= '0'+9 {
		return num
	}
	return nonNum
}
