func findMin(nums []int) int {
	res := nums[0]

	l, r := 0, len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] < res {
			res = nums[mid]
		}

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return res
}
