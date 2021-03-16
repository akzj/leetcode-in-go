package binary_tree_level_order_traversal_102

import "github.com/akzj/leetcode-in-go/structures"

func LevelOrderTraversal(root *structures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var nodes = []*structures.TreeNode{root}
	for len(nodes) != 0 {
		temp := nodes
		nodes = nil
		var vals []int
		for _, node := range temp {
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
			vals = append(vals, node.Val)
		}
		result = append(result, vals)
	}
	return result
}
