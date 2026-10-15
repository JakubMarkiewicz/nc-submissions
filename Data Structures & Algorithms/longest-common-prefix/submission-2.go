func longestCommonPrefix(strs []string) string {
	res := ""

	for {
		idx := len(res)
		same := true
		for _, value := range strs {
			overflow := idx >= len(value)
			if overflow || value[idx] != strs[0][idx] {
				same = false
				break
			}
			
		}

		if !same {
			break
		} else {
			res = res + string(strs[0][idx])
		}
	}

	return res
}
