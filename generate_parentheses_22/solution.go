package generate_parentheses_22

import "strings"

func GenerateParentheses(n int) []string {
	return exploreParentheses(0, n, nil)
}

func exploreParentheses(count, maxCount int, mem []string) []string {
	if len(mem) == maxCount*2 {
		if count == 0 {
			return []string{strings.Join(mem, "")}
		}
		return nil
	}
	if count == 0 {
		return exploreParentheses(count+1, maxCount, append(mem, "("))
	} else if count > 0 && count < maxCount {
		var result []string
		if res := exploreParentheses(count+1, maxCount, append(mem, "(")); res != nil {
			result = append(result, res...)
		}
		if res := exploreParentheses(count-1, maxCount, append(mem, ")")); res != nil {
			result = append(result, res...)
		}
		return result
	} else {
		return exploreParentheses(count-1, maxCount, append(mem, ")"))
	}
}
