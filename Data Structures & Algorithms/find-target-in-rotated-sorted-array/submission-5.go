func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2	

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
		// sorted to left
			if target >= nums[l] && target < nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}


		} else {
		// sorted to right
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		}

	}

	if nums[l] == target {
		return l
	}

	return -1
}
