func characterReplacement(s string, k int) int {
	chars := make(map[byte]bool)

	for idx := range s {
		chars[s[idx]] = true
	}

	res := 0

	for b := range chars {
		matchCount, l := 0,0		
		for r := 0; r < len(s); r++ {
			if s[r] == b {
				matchCount++
			}
			
			for  r-l + 1 - matchCount > k {
				if s[l] == b {
					matchCount--
				}
				l++
			}
			res = max(res, r - l + 1)

		}
	}

	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
