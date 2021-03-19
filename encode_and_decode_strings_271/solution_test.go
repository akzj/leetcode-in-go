package encode_and_decode_strings_271

import (
	"fmt"
	"testing"
)

func TestCodec(t *testing.T) {
	coder := Codec{}

	fmt.Println(coder.Decode(coder.Encode([]string{"hello", "world", "sxxxxxxxxxxxxxx", "jjjjjjjjj"})))
}
