func isValid(s string) bool {
	matches := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, v := range s {
		pair, isClosing := matches[v]

		if isClosing {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if last != pair {
				return false
			}
			stack = stack[0:len(stack)-1]
		} else {
			stack = append(stack,v)
		}

	}

return len(stack) == 0
}
