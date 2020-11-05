package main

import "fmt"

// https://leetcode-cn.com/problems/reverse-linked-list/

var testData *ListNode

func init() {
    testData = &ListNode{Val: 1}

    forData := testData
    for i := 2; i <= 5; i++ {
        forData.Next = &ListNode{Val: i}
        forData = forData.Next
    }
}

func main() {
    printList(testData)
    printList(reverseList1(testData))
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func printList(node *ListNode)  {
    for  {
        if node == nil {
            break
        }
        fmt.Println(node.Val)
        node = node.Next
    }
}

// 不需要数组，直接原链表改造
func reverseList1(head *ListNode) *ListNode {
    var pre *ListNode
    var cur = head
    for {
        if cur == nil {
            break
        }
        tmp := cur.Next // next 缓存，下面要被覆盖掉
        cur.Next = pre // 当前的 next 赋值成 循环的上一个
        pre = cur // 上一个就是当前循环的 cur，第一个次循环 pre 是 nil
        cur = tmp // 当前的就是 当前的 next
    }
    return pre
}

// 最简单方法，数组方式
// 空间复杂度 O(n)
// 但是速度会很快不知道为啥
func reverseList(head *ListNode) *ListNode {
    var s []int
    for {
        if head == nil {
            break
        }
        s = append(s, head.Val)
        head = head.Next
    }
    var retHead *ListNode
    var ret *ListNode
    for i := len(s) - 1; i >= 0; i-- {
        if ret == nil {
            ret = &ListNode{Val: s[i]}
            retHead = ret
        } else {
            ret.Next = &ListNode{Val: s[i]}
            ret = ret.Next
        }
    }
    return retHead
}
