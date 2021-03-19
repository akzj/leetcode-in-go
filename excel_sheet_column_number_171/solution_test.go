package excel_sheet_column_number_171

import (
	"fmt"
	"testing"
)

func TestGetColumnNumber(t *testing.T) {

	fmt.Println(GetColumnNumber("A"))
	fmt.Println(GetColumnNumber("AA"))
	fmt.Println(GetColumnNumber("ZY"))
}
