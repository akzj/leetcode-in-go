package reverse_string_344

func ReverseString(data []byte) {
	var i = 0
	var j = len(data) - 1
	for ; i < j; {
		c := data[i]
		data[i] = data[j]
		data[j] = c
		i++
		j--
	}
}
