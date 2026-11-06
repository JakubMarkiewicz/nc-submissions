func characterReplacement(s string, k int) int {
	chars := make(map[byte]struct{})	

	for idx := range s {
		chars[s[idx]] = struct{}{}
	}

	res := 0

	for char := range chars {
		l, r := 0, 0	
		matchCount := 0
		for r < len(s) {
			if s[r] == char {
				matchCount++
			}
			for (r-l + 1) - matchCount > k {
				if s[l] == char {
					matchCount--
				}
				l++
			}
			if (r-l + 1) > res {
				res = r-l + 1
			}
			r++	
		}
	}
	
	return res
}
