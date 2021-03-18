package decrypt_string_from_alphabet_to_integer_mapping_1309

import (
	"fmt"
	"strconv"
)

func Decrypt(data string) string {

	var decodeMap = map[string]string{}

	for i := 1; i <= 9; i++ {
		decodeMap[strconv.Itoa(i)] = string('a' + byte(i-1))
	}

	for i := 10; i <= 26; i++ {
		decodeMap[reverseString([]byte(strconv.Itoa(i)+"#"))] = string('j' + byte(i-10))
	}

	var result string
	data = reverseString([]byte(data))
	for len(data) > 0 {
		ch := data[0]
		switch ch {
		case '#':
			key := data[:3]
			fmt.Println(key, decodeMap[key])
			result += decodeMap[key]
			data = data[3:]
		default:
			result += decodeMap[string(ch)]
			data = data[1:]
		}
	}
	return reverseString([]byte(result))
}

func reverseString(data []byte) string {
	var i = 0
	var j = len(data) - 1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return string(data)
}

func Decrypt2(data string) string {
	var table [27]int32
	for i := 1; i <= 26; i++ {
		table[i] = 'a' + int32(i-1)
	}

	var result string
	for i := 0; i < len(data); {
		if len(data) >= 3 && data[i+2] == '#' {
			index, _ := strconv.Atoi(data[i : i+2])
			result += string(table[index])
			i += 3
		} else {
			result += string(table[data[i]-'1'])
			i += 1
		}
	}
	return result
}
