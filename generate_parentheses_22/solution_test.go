package generate_parentheses_22

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(strings.Join(GenerateParentheses(3),"\n"))
	fmt.Println(strings.Join(GenerateParentheses(5),"\n"))
}
