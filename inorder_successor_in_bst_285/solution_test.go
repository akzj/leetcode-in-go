package inorder_successor_in_bst_285

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestGetSuccessor(t *testing.T) {

	if successor := GetSuccessor(structures.BuildTree([]int{2, 1, 3}), 1); successor != nil {
		fmt.Println(successor.Val)
	}

	if successor := GetSuccessor(structures.BuildTree([]int{5,3,6,2,4,1}), 5); successor != nil {
		fmt.Println(successor.Val)
	}else{
		fmt.Println("null")
	}

	if successor := GetSuccessor(structures.BuildTree([]int{5,3,6,2,4,1}), 6); successor != nil {
		fmt.Println(successor.Val)
	}else{
		fmt.Println("null")
	}
}
