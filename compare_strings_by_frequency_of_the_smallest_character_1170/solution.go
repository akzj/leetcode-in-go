package compare_strings_by_frequency_of_the_smallest_character_1170

import "math"

func CompareStrings(queries []string, words []string) []int {

	var frequencyMap = make(map[string]int)

	getFrequency := func(str string) int {
		qf, ok := frequencyMap[str]
		if ok == false {
			qf = frequencySmallestCharacter(str)
			frequencyMap[str] = qf
		}
		return qf
	}

	var result = make([]int, len(queries))
	for index, query := range queries {
		for _, word := range words {
			if getFrequency(query) < getFrequency(word) {
				result[index]++
			}
		}
	}
	return result
}

func frequencySmallestCharacter(str string) int {
	var smallC = int32(math.MaxInt32)
	var frequency int
	for _, c := range str {
		if c < smallC {
			frequency = 1
			smallC = c
		} else if smallC == c {
			frequency++
		}
	}
	return frequency
}
