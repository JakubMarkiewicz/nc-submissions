func checkInclusion(s1 string, s2 string) bool {
	set := make(map[byte]struct{})
	s1Counts := make(map[byte]int)

	for idx := range s1 {
		set[s1[idx]] = struct{}{}
		s1Counts[s1[idx]]++
	}

	l, r := 0,0

	for r < len(s2) {
		if _, ok := set[s2[r]]; !ok {
			for l < r {
				if _, ok := set[s2[l]]; ok {
					s1Counts[s2[l]]++
				}
				l++
			}
			l++
			r = l
			continue
		}

		if c := s1Counts[s2[r]]; c > 0 {
			s1Counts[s2[r]]--
			r++
		} else {
			for l < r {
				if _, ok := set[s2[l]]; ok {
					s1Counts[s2[l]]++
				}
				if s2[l] == s2[r] {
					l++
					break
				}
				l++
			}
		}

		if r-l == len(s1) {
			return true
		}


	}

	return false
}
