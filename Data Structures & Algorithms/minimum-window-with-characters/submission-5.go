    func minWindow(s string, t string) string {

       sCountTmp := make(map[byte]int)

       for idx := range t {
          sCountTmp[t[idx]]++
       }

       missing := len(t)

       res := []int{-1,-1}

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
             if res[0] == -1 ||  r-l < res[1] - res[0] {
                res = []int{l,r}
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

	   if res[0] == -1 {
		return ""
	   }

       return s[res[0]:res[1]+1]
    }