func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	l1, l2 := 0, 0

	for l1 < len(word1) || l2 < len(word2) {
		if l1 < len(word1) {
			res.WriteByte(word1[l1])
			l1++
		}
		if l2 < len(word2) {
			res.WriteByte(word2[l2])
			l2++
		}
	}

	return res.String()
}
