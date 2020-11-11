package main

import "shijianzhiwai/thinking-leetcode/lib"

//https://leetcode-cn.com/problems/invert-binary-tree/

var testData *lib.TreeNode

func init() {
	testData = &lib.TreeNode{
		Val: 4,
		Left: &lib.TreeNode{
			Val: 2,
			Left: &lib.TreeNode{
				Val: 1,
			},
			Right: &lib.TreeNode{
				Val: 3,
			},
		},
		Right: &lib.TreeNode{
			Val: 7,
			Left: &lib.TreeNode{
				Val: 6,
			},
			Right: &lib.TreeNode{
				Val: 9,
			},
		},
	}
}

func main() {
	lib.PrintTree(invertTree(testData))
}

func invertTree(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
