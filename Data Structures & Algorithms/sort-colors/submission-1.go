func sortColors(nums []int) {
	for i := 1; i < len(nums); i++ {
		currentNum := nums[i]

		if currentNum < nums[i-1] {
			nums[i] = nums[i-1]
			nums[i-1] = currentNum	
			if i == 1 {
				i -= 1
			} else {
				i -= 2
			}
			continue
		}
	}
}
