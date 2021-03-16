package reverse_string_ii_541

import (
	"fmt"
	"testing"
)

func TestReverseStringK(t *testing.T) {


	var data = []byte("abcdefg")

	ReverseStringK(data,2)

	fmt.Println(string(data))
}
