package binary_tree_maximum_path_sum_124

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestBinaryTreeMaximumPathSum(t *testing.T) {

	fmt.Println(pathSum(&structures.TreeNode{
		Left: &structures.TreeNode{
			Val: 9,
		},
		Right: &structures.TreeNode{
			Left: &structures.TreeNode{
				Val: 15,
			},
			Right: &structures.TreeNode{
				Val: 7,
			},
			Val: 20,
		},
		Val: -10,
	}))
}
