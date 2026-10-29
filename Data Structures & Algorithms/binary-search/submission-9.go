func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l < r {
		middle := l + (r-l) / 2

		num := nums[middle]

		 if num >= target {
			r = middle
		} else {
			l = middle + 1
		}
		
	}

	if nums[l] == target {
		return l
	} 

	return -1
}
