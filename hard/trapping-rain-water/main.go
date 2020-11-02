package main

import "fmt"

// https://leetcode-cn.com/problems/trapping-rain-water/

/*
## 双指针法

左右俩指针，左右同时向中间移动
----

> 认为如果一端有更高的条形块（例如右端），积水的高度依赖于当前方向的高度（从左到右）。当我们发现另一侧（右侧）的条形块高度不是最高的，我们则开始从相反的方向遍历（从右到左）。
作者：LeetCode
链接：https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

var testData = []int{0,1,0,2,1,0,1,3,2,1,2,1}

func main() {
    fmt.Println(trap(testData))
}

func trap(height []int) int {
    var trap int
    var left, right int
    right = len(height) - 1
    var leftMax, rightMax int
    for left < right {
        if height[left] < height[right] {
            if leftMax <= height[left] {
                // 如果当前的值比之前的大，那么之前的循环一定是无法存贮水的，所以不会计算水，只会继续向右移动
                // 注意这个等于号，等于的时候也是无法存储水的
                leftMax = height[left]
            } else {
                // 这个时候之所以可以存储水，右边呢？别忘了右边在这个循环里面一定是大于当前的，所以必然是可以存储水的
                trap += leftMax-height[left]
                // 证明 leftMax <= height[right]
                // 当前 if 是只有 height[left] < height[right] 的时候才会进入，所以当前 if 产生的 leftMax 也一定是小于 height[right]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                trap += rightMax - height[right]
            }
            right--
        }
    }
    return trap
}