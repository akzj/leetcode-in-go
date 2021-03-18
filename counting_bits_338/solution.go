package counting_bits_338

func CountingBits(num int) []int {
	var bits = make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

func CountingBits3(num int) []int {
	var bits = make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

func CountingBits2(num int) []int {
	var bits = make([]int, num+1)
	bits[0] = 0
	for i := 1; i <= num; i++ {
		bits[i] = getBitsCount(i)
	}

	return bits
}

func getBitsCount(num int) int {
	var n int
	for num > 0 {
		num &= num - 1
		n++
	}
	return n
}
