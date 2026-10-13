class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    validPalindrome(s) {
      let r = s.length - 1
      let l = 0
      let changed = false
      
      while (l < r) {
        if (s[l] === s[r]) {
            l++
            r--
            continue
        } else if (changed) {
            return false
        }
        if (s[l] === s[r-1]) {
            r--
        } else {
            l++
        }
        changed = true
      }

        return true
    }
}
