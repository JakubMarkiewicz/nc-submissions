func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]int)
	maxL := 0
	l, r := 0, 0

	for r < len(s) {
		if idx, ok := set[s[r]]; ok {
			l = max(idx+1, l)
		}	
		set[s[r]] = r
		maxL = max(maxL, r-l+1)
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

