package construct_binary_tree_from_preorder_and_inorder_traversal_105

import "github.com/akzj/leetcode-in-go/structures"

func BuildTree(preorder, inorder []int) *structures.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var node = &structures.TreeNode{Val: preorder[0]}
	index := getIndex(inorder, preorder[0])

	preorder = preorder[1:]
	node.Left = BuildTree(preorder[:len(inorder[:index])], inorder[:index])
	node.Right = BuildTree(preorder[len(inorder[:index]):], inorder[index+1:])
	return node
}

func getIndex(vals []int, target int) int {
	for i, val := range vals {
		if val == target {
			return i
		}
	}
	return -1
}
