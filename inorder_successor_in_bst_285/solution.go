package inorder_successor_in_bst_285

import "github.com/akzj/leetcode-in-go/structures"

func GetSuccessor(node *structures.TreeNode, key int) *structures.TreeNode {
	var successor *structures.TreeNode

	for node != nil {
		if node.Val > key {
			if successor == nil || node.Val < successor.Val {
				successor = node
			}
		}
		if node.Val > key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}
