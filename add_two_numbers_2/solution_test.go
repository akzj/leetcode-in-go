package add_two_numbers_2

import (
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestName(t *testing.T) {

	sumList := AddTwoNumbers(structures.BuildList([]int{2, 4, 3}), structures.BuildList([]int{5, 6, 4}))

	structures.PrintList(sumList)
}
