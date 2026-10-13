func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]int)

	for i  := range s {
		sMap[s[i]] += 1
		sMap[t[i]] -= 1
	}

	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}

	return true
}
