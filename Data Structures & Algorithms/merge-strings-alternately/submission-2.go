func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	l := 0

	for l < len(word1) || l < len(word2) {
		if l < len(word1) {
			res.WriteByte(word1[l])
		}
		if l < len(word2) {
			res.WriteByte(word2[l])
		}
		l++
	}

	return res.String()
}
