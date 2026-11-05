func lengthOfLongestSubstring(s string) int {

	set := make(map[byte]bool)
	maxL := 0
	l, r := 0, 0

	for r < len(s) {
		for set[s[r]] {
			delete(set, s[l])	
			l++
		}

		set[s[r]] = true
		maxL = max(maxL, r-l + 1)

		r++
	}

	return maxL
}

func max(a,b int) int {
	if a > b {
		return a 
	}
	return b
}

