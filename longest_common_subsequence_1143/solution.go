package longest_common_subsequence_1143

func LongCommonSubSequence(text1, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}

	charIndex := make(map[int32][]int)

	for index, c := range text1 {
		charIndex[c] = append(charIndex[c], index)
	}

	var next int
	for i, c := range text2 {
		indexes, ok := charIndex[c]
		if ok == false {
			return i
		}
		var match = false
		for _, index := range indexes {
			if index < next {
				continue
			}
			next = index
			match = true
			break
		}
		if match == false {
			return i
		}
	}
	return len(text2)
}
