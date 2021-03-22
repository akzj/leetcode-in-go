package fizz_buzz_412

import "strconv"

func FizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		var str string
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str += strconv.Itoa(i)
		}
		result = append(result, str)
	}
	return result
}
