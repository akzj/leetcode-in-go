package letter_combinations_of_a_phone_number_17

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	mapping := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	combs := mapping[string(digits[0])]
	digits = digits[1:]
	if len(digits) == 0 {
		return combs
	}

	for _, c := range digits {
		var temp []string
		nextCombs := mapping[string(c)]
		for _, str := range combs {
			for _, str2 := range nextCombs {
				temp = append(temp, str+str2)
			}
		}
		combs = temp
	}
	return combs

}
