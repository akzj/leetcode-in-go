package longest_common_prefix_14

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
}
