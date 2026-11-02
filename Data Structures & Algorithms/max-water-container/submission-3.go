func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	max := 0

	l, r := 0, len(heights) - 1

	for l < r {
		area := (r-l) * min(heights[l], heights[r])

		if area > max {
			max = area
		}

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}

	}

	return max
}
