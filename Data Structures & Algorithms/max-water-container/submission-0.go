func getMin(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func maxArea(heights []int) int {
	max := 0

	l, r := 0, len(heights) - 1

	for l < r {
		localMin := getMin(heights[l], heights[r])
		width := r-l

		if newMax := width*localMin; newMax > max {
			fmt.Printf("l: %d, r: %d\n", l,r)
			max = newMax	
		}

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}

	}

	return max
}
