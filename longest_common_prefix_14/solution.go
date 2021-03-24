package longest_common_prefix_14

import "math"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var index = math.MaxInt32
	for i, str := range strs {
		if len(str) < index {
			index = i
		}
	}
	var prefix string
	for i := range strs[index] {
		var match = true
		for j, str := range strs {
			if j == index {
				continue
			}
			if str[i] == strs[index][i] {
				continue
			}
			match = false
			break
		}

		if match == false {
			return prefix
		}
		prefix += string(strs[index][i])
	}
	return prefix
}
