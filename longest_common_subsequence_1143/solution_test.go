package longest_common_subsequence_1143

import (
	"fmt"
	"testing"
)

func TestLongCommonSubSequence(t *testing.T) {

	fmt.Println(LongCommonSubSequence("abcde","ace"))
	fmt.Println(LongCommonSubSequence("abc","abc"))
	fmt.Println(LongCommonSubSequence("123","456"))
	fmt.Println(LongCommonSubSequence("123","415263"))
	fmt.Println(LongCommonSubSequence("123","415326"))
}
