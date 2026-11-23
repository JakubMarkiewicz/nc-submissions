// modify input array, go to idx 
func findDuplicate(nums []int) int {
	idx := nums[0]

	for true {
		if nums[idx] == -1 {
			return idx
		}
		newIdx := nums[idx]
		nums[idx] = -1
		idx = newIdx
	}

	return -1
}
