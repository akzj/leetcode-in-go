package reverse_string_ii_541

func ReverseStringK(data []byte, k int) {
	var begin int
	for begin < len(data) {
		var i = begin
		var j = begin + k - 1
		for i < j && j < len(data)-1 && i < len(data)-1 {
			c := data[i]
			data[i] = data[j]
			data[j] = c
			i++
			j--
		}
		begin += 2 * k
	}
}
