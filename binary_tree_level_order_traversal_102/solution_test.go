package binary_tree_level_order_traversal_102

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {

	root := structures.BuildTree([]int{10, 5, 3, 7, 4, 1, 15, 11, 16, 12, 13})

	structures.PrintTree(root,1)

	fmt.Println(LevelOrderTraversal(root))
}
