func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l < r {
		middle := (l + r) / 2	

		num := nums[middle]

		if l == middle || r == middle {
			break
		}
		 if num > target {
			r = middle - 1
		} else {
			l = middle
		}
		
	}

	if nums[l] == target {
		return l
	} else if nums[r] == target {
		return r
	}

	return -1
}
