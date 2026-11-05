func trap(height []int) int {
	l, r := 0, len(height)-1
	maxL, maxR := 0, 0

	res := 0

	for l < r {	
			maxL = max(maxL, height[l])
			maxR = max(maxR, height[r])
		if maxL < maxR {
			res += max(min(maxL, maxR) - height[l], 0)
			l++
		} else {
			res += max(min(maxL, maxR) - height[r], 0)
			r--
		}
	}

	return res
}


func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
