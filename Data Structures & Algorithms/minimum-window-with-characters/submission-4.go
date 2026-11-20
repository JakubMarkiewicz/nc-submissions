    func minWindow(s string, t string) string {
       sCountTmp := make(map[byte]int)

       for idx := range t {
          sCountTmp[t[idx]]++
       }

       missing := len(t)

       res := ""

       l, r := 0, 0
       for r < len(s) {
          if _, ok := sCountTmp[s[l]]; l == r && !ok {
             l++
             r = l
             continue
          }

          rChar := s[r]

		if _, ok := sCountTmp[rChar]; ok {
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
			if _, ok := sCountTmp[lChar]; ok {
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