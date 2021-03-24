package longest_increasing_subsequence_300

import (
	"fmt"
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {

	fmt.Println(LongestIncreasingSubsequence([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(LongestIncreasingSubsequence2([]int{10,9,2,5,3,7,101,18}))

	fmt.Println(LongestIncreasingSubsequence([]int{1,2,3}))
	fmt.Println(LongestIncreasingSubsequence2([]int{1,2,3}))



	fmt.Println(LongestIncreasingSubsequence([]int{3,2,1}))
	fmt.Println(LongestIncreasingSubsequence2([]int{3,2,1}))
}
