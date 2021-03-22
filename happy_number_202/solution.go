package happy_number_202

func HappyNumber(num int) bool {
	var nextMap = make(map[int]bool)
	var next = num
	for {
		next = nextNum(next)
		if next == 1 {
			return true
		}
		if nextMap[next] {
			return false
		}
		nextMap[next] = true
	}
}

func nextNum(num int) int {
	var next int
	for num != 0 {
		next += (num % 10) * (num % 10)
		num = num / 10
	}
	return next
}
