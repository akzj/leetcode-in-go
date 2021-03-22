package first_unique_character_in_a_string_387

func findUniqueCharacter(str string) int {
	var characters = map[int32]int{}
	for _, ch := range str {
		characters[ch] += 1
	}
	for index, ch := range str {
		if count, _ := characters[ch]; count == 1 {
			return index
		}
	}
	return -1
}
