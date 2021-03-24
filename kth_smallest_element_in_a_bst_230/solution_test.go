package kth_smallest_element_in_a_bst_230

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestKthSmallestElement(t *testing.T) {

	fmt.Println(KthSmallestElement(structures.BuildTree([]int{5,3,6,2,4,1}), 3))
	fmt.Println(KthSmallestElement(structures.BuildTree([]int{5,3,6,2,4,1}), 6))
	fmt.Println(KthSmallestElement(structures.BuildTree([]int{3,1,4,2}), 1))
}
