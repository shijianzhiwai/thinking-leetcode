package main

import "shijianzhiwai/thinking-leetcode/lib"

// https://leetcode-cn.com/problems/merge-two-sorted-lists/

var l1 *lib.ListNode
var l2 *lib.ListNode

func init() {
	l1 = &lib.ListNode{Val: 1, Next: &lib.ListNode{Val: 2, Next: &lib.ListNode{Val: 5}}}
	l2 = &lib.ListNode{Val: 1, Next: &lib.ListNode{Val: 3, Next: &lib.ListNode{Val: 4}}}
}

func main() {
	//lib.PrintList(l1)
	//lib.PrintList(l2)
	lib.PrintList(mergeTwoLists(l1,l2))
}

func mergeTwoLists(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var ret *lib.ListNode
	var cur *lib.ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			if ret == nil {
				ret = l1
				cur = l1
			} else {
				tmp := cur
				cur = l1
				tmp.Next = cur
			}
			l1 = l1.Next
		} else {
			if ret == nil {
				ret = l2
				cur = l2
			} else {
				tmp := cur
				cur = l2
				tmp.Next = cur
			}
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return ret
}
