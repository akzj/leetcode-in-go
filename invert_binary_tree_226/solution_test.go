package invert_binary_tree_226

import (
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {

	structures.PrintTree(InvertBinaryTree(structures.BuildTree([]int{4, 2, 1, 3, 7, 6, 9})), 1)
}
