package excel_sheet_column_number_171

func GetColumnNumber(title string) int {
	var sum int
	for _, ch := range title {
		num := ch - 'A' + 1
		sum = sum*26 + int(num)
	}
	return sum
}
