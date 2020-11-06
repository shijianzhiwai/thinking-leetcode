package lib

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode)  {
	if root == nil {
		return
	}
	PrintTree(root.Left)
	PrintTree(root.Right)

	fmt.Println(root.Val)
}