package group_anagrams_49




func GroupAnagrams(strs []string) [][]string {
	var groups [][]string
	var remain = strs
	for remain != nil {
		var group = []string{remain[0]}
		var temp []string
		for _, str := range remain[1:] {
			if IsAnagrams(group[0], str) {
				group = append(group, str)
			} else {
				temp = append(temp, str)
			}
		}
		groups = append(groups, group)
		remain = temp
	}
	return groups
}

func IsAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var chars = map[int32]int{}
	for _, c := range a {
		chars[c] += 1
	}
	for _, c := range b {
		if v, ok := chars[c]; ok && v > 0 {
			chars[c] -= 1
			continue
		}
		return false
	}

	for _, value := range chars {
		if value != 0 {
			return false
		}
	}

	return true
}
