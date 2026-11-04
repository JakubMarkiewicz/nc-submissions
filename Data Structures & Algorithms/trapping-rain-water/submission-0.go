func intMax(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func intMin(a,b int) int {
	if a < b {
		return a
	}

	return b
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	max := 0

	fromLeft := []int{}
	fromRight := []int{}

	localHighestBar := 0
	for idx := 0; idx < len(height); idx++ {
		barH := height[idx]
		localHighestBar = intMax(barH, localHighestBar)

		if idx != 0 {
			fromLeft = append(fromLeft, localHighestBar - barH)
		} else {
			fromLeft = append(fromLeft, 0)
		}
	}
	localHighestBar = 0
	for idx := len(height) - 1; idx >= 0; idx-- {
		barH := height[idx]
		localHighestBar = intMax(barH, localHighestBar)

		if idx !=0 {
		fromRight = append(fromRight, localHighestBar - barH)
		} else {
		fromRight = append(fromRight, 0)
		}
	}

	for idx := range height {
		max += intMin(fromLeft[idx], fromRight[len(height) - idx -1])
	}




	return max
}
