package reverse_string_344

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {

	var data = []byte("1234")
	ReverseString(data)
	fmt.Println(string(data))
}
