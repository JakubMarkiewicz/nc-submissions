func sortColors(nums []int) {
	for i := 1; i < len(nums); i++ {
		currentNum := nums[i]

		if currentNum < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i -= 2
			if i < 0 {
				i = 0
			}
			continue
		}
	}
}
