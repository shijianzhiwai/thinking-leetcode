package main

import (
	"fmt"
	"shijianzhiwai/thinking-leetcode/lib"
)

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
//    1
///   \
//  2   5
///  \   \
//  3   4   6
var testData *lib.TreeNode

func init() {
	testData = &lib.TreeNode{
		Val: 1,
		Left: &lib.TreeNode{
			Val: 2,
			Left: &lib.TreeNode{
				Val: 3,
			},
			Right: &lib.TreeNode{
				Val: 4,
			},
		},
		Right: &lib.TreeNode{
			Val: 5,
			Right: &lib.TreeNode{
				Val: 6,
			},
		},
	}
}

func main() {
	flatten(testData)
}

func flatten(root *lib.TreeNode) {
	if root == nil {
		return
	}
	arr := make([]int, 0)
	getList(root, &arr)
	root.Left = nil
	tmp := root
	fmt.Println(arr)
	for k,v := range arr{
		if k == 0 {
			continue
		}
		tmp.Right = &lib.TreeNode{Val: v}
		tmp = tmp.Right
	}
}

func getList(root *lib.TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	*arr = append(*arr, root.Val)
	getList(root.Left, arr)
	getList(root.Right, arr)
}
