package main

import "fmt"

// https://leetcode-cn.com/problems/trapping-rain-water/

/*
## 双指针法

左右俩指针，左右同时向中间移动
    * 水的高度取决于低的一边，当一遍不再是最低的时候，就开始反向遍历
    * 所以那边低，哪边向中间移动，直到他不是最低的
    * 同时左右移动边移动边保存两边指针最高的，最高的就把之前较为矮的挡住，所以谁存多少取决于之后的高的

左右指针循环为什么行得通？
【左边的某个位置之后的位置高度，如果右边出现等于或者高于这个高度的，则这个位置必然可以存水】

----

> 以认为如果一端有更高的条形块（例如右端），积水的高度依赖于当前方向的高度（从左到右）。当我们发现另一侧（右侧）的条形块高度不是最高的，我们则开始从相反的方向遍历（从右到左）。
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
                leftMax = height[left]
            } else {
                trap += leftMax-height[left]
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