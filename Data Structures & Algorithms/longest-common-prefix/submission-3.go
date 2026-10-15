func longestCommonPrefix(strs []string) string {
	var res strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		for _, value := range strs {
			overflow := i >= len(value)
			if overflow || value[i] != strs[0][i] {
				return res.String()
			}
			
		}

res.WriteByte(strs[0][i])
	}

	return res.String()
}
