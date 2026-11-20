    func minWindow(s string, t string) string {
       sCount := make(map[byte]int)
       sCountTmp := make(map[byte]int)

       for idx := range t {
          sCount[t[idx]]++
          sCountTmp[t[idx]]++
       }

       missing := len(t)

       res := ""

       l, r := 0, 0
       for r < len(s) {
          if _, ok := sCount[s[l]]; l == r && !ok {
             l++
             r = l
             continue
          }

          rChar := s[r]

		if _, ok := sCount[rChar]; ok {
			if sCountTmp[rChar] > 0 {
				missing--
			}
			sCountTmp[rChar]--
		}


          for missing == 0 {
             if res == "" || r-l < len(res) {
                res = s[l:r+1]
             }
			lChar := s[l]
			if _, ok := sCount[lChar]; ok {
				sCountTmp[lChar]++
				if sCountTmp[lChar] > 0 {
					missing++
				}
			}

             l++
          }

          r++
       }

       return res
    }