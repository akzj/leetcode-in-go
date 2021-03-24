package kth_smallest_element_in_a_bst_230

import "github.com/akzj/leetcode-in-go/structures"

func KthSmallestElement(node *structures.TreeNode, k int) int {
	return getKthElement(node, &k).Val
}

func getKthElement(node *structures.TreeNode, k *int) *structures.TreeNode {
	if node == nil {
		return nil
	}
	left := getKthElement(node.Left, k)
	if left != nil {
		return left
	}
	*k = *k - 1
	if *k == 0 {
		return node
	}
	return getKthElement(node.Right, k)
}
