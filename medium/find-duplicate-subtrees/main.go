package main

import (
	"fmt"
	"shijianzhiwai/thinking-leetcode/lib"
)

// https://leetcode-cn.com/problems/find-duplicate-subtrees/

var testData *lib.TreeNode

func init() {
	testData = &lib.TreeNode{
		Val: 1,
		Left: &lib.TreeNode{
			Val:  2,
			Left: &lib.TreeNode{Val: 4},
		},
		Right: &lib.TreeNode{
			Val: 3,
			Left: &lib.TreeNode{
				Val: 2,
				Left: &lib.TreeNode{
					Val: 4,
				},
			},
			Right: &lib.TreeNode{
				Val: 4,
			},
		},
	}
}

var ret []*lib.TreeNode
var mapData map[string]int

func main() {
	resp := findDuplicateSubtrees(testData)
	for _, r := range resp {
		lib.PrintTree(r)
	}
	fmt.Println("count", count(testData))

	fmt.Println("======")
	traverse1(testData)
	fmt.Println("======")
	traverse2(testData)
	fmt.Println("======")
	traverse3(testData)
	fmt.Println("======")
}

func findDuplicateSubtrees(root *lib.TreeNode) []*lib.TreeNode {
	mapData = make(map[string]int)
	ret = make([]*lib.TreeNode, 0)
	traverse(root)
	return ret
}

func traverse(root *lib.TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left)
	right := traverse(root.Right)

	str := left + "," + right + "," + fmt.Sprintf("%d", root.Val)
	resp := mapData[str]
	if resp == 0 {
		mapData[str] = 1
	} else {
		c := resp + 1
		if c == 2 {
			ret = append(ret, root)
		}
		mapData[str] = c
	}
	return str
}

func count(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}

	left := count(root.Left)
	right := count(root.Right)

	return left + right + 1
}

func traverse1(root *lib.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	traverse1(root.Left)
	traverse1(root.Right)
}

func traverse2(root *lib.TreeNode) {
	if root == nil {
		return
	}
	traverse2(root.Left)
	fmt.Println(root.Val)
	traverse2(root.Right)
}

func traverse3(root *lib.TreeNode) {
	if root == nil {
		return
	}
	traverse3(root.Left)
	traverse3(root.Right)
	fmt.Println(root.Val)
}
