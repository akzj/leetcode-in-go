package convert_integer_to_the_sum_of_two_no_zero_integers

func GetNoZeroIntegers(n int) []int {

	var one = n
	var two = 0

	for hasZero(one) || hasZero(two) {
		one--
		two++
	}

	return []int{one, two}
}

func hasZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}
