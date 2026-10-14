func groupAnagrams(strs []string)(result[][]string) {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		key := [26]int{}
		for _, b := range str {
			key[b-'a'] += 1
		}
		m[key] = append(m[key], str)
	}

	for _, v := range m {
		result = append(result,v)
	}

	return result


}
