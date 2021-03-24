package invert_binary_tree_226

import "github.com/akzj/leetcode-in-go/structures"

func InvertBinaryTree(head *structures.TreeNode) *structures.TreeNode {
	if head == nil {
		return nil
	}
	head.Left, head.Right = InvertBinaryTree(head.Right), InvertBinaryTree(head.Left)
	return head
}
