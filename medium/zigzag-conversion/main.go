package main

import "fmt"

// https://leetcode-cn.com/problems/zigzag-conversion/

func main() {
    fmt.Println(convert("LEETCODEISHIRING", 3))
    fmt.Println(convert("LEETCODEISHIRING", 4))
}

// 第一种解法
func convert(s string, numRows int) string {
    count := numRows - 2
    pool := make([][]rune, numRows)
    rs := []rune(s)
    index := 0
    bigFor := 0
    // 这个函数只是打印 rune 数组
    retFn := func() string {
        var ret []rune
        for _, p := range pool {
            for _, v := range p {
                ret = append(ret, v)
            }
        }
        return string(ret)
    }
    for {
        // bigFor 控制方向
        if bigFor == 0 {
            for k := range pool {
                // 循环行数，每一行都去索取一个字母
                if index == len(s) {
                    return retFn()
                }
                pool[k] = append(pool[k], rs[index])
                index++
            }
            // 索取完毕，转换方向
            bigFor++
        } else {
            // 每一个完整竖行之间的数量为 count = numRows-2
            // 从最下面开始向上循环，每次 pool[i] 向上移动一行，并索取一个字母
            for i := count; i >= 1; i-- {
                if index == len(s) {
                    return retFn()
                }
                pool[i] = append(pool[i], rs[index])
                index++
            }
            // 索取完毕，转换方向
            bigFor = 0
        }
    }
}

// 第二种解法
func convert1(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    pool := make([][]byte, numRows)
    flag := -1
    row := 0
    for _, v := range s {
        // 当循环到底或者到顶，转换方向
        // flag 赋值相反，等于翻转方向
        // 将 v 填进去到对用的 pool 行数
        if row == 0 || row == numRows-1 {
            flag = -flag
        }

        pool[row] = append(pool[row], byte(v))
        row += flag
    }
    ret := ""
    for _,p := range pool{
        ret+=string(p)
    }
    return ret
}