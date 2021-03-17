package binary_tree_maximum_path_sum_124

import (
	"github.com/akzj/leetcode-in-go/structures"
	"math"
)

var max int

func pathSum(node *structures.TreeNode)int {
	explorePathSum(node)
	return max
}

func explorePathSum(head *structures.TreeNode) int {
	if head == nil {
		return 0
	}
	leftSum := math.Max(0.0, float64(explorePathSum(head.Left)))
	rightSum := math.Max(0.0, float64(explorePathSum(head.Right)))
	pathSum := head.Val + int(leftSum+rightSum)
	max = int(math.Max(float64(max), float64(pathSum)))

	return head.Val + int(math.Max(leftSum, rightSum))
}
