package construct_binary_tree_from_preorder_and_inorder_traversal_105

import (
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestBuildTree(t *testing.T) {
	structures.PrintTree(BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}), 0)
}
