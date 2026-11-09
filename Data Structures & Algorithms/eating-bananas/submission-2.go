
func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)

	l,r := 1, piles[len(piles)-1]


	min := -1
	for l <= r {
		total := 0
		mid := l + (r-l)/2

		for idx := range piles {
			hoursNeeded := int(math.Ceil(float64(piles[idx]) / float64(mid)))
			total += hoursNeeded
		}	


		if total > h {
			l = mid + 1
			continue
		}

		if min == -1 || mid < min {
			min = mid
		}
		r = mid -1
		
	}

	return min
	
}
