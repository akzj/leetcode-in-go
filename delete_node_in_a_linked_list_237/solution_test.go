package delete_node_in_a_linked_list_237

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestDeleteNode(t *testing.T) {

	structures.PrintList(DeleteNode(structures.BuildList([]int{1, 2, 3, 4, 5, 6}), 5))
	fmt.Println("")
	structures.PrintList(DeleteNode(structures.BuildList([]int{1, 2, 3, 4, 5, 6}), 6))
	fmt.Println("")
	structures.PrintList(DeleteNode(structures.BuildList([]int{1, 2, 3, 4, 5, 6}), 1))
	fmt.Println("")
}
