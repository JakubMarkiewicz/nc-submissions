func nextGreaterElement(nums1 []int, nums2 []int) []int { 
	m := make(map[int]int)

	for i, v := range nums1 {
		m[v] = i
	}

	res := make([]int, len(nums1))	

	for i := range res {
		res[i] = -1
	}

	stack := []int{}	

	for _, value := range nums2 {
		for len(stack) > 0 && value > stack[len(stack) - 1]  {
			lastVal := stack[len(stack) - 1]
			resIdx := m[lastVal]	
			res[resIdx] = value
			stack = stack[0 : len(stack) - 1]
		}

		if _, ok := m[value]; ok {
			stack = append(stack, value)
		}
	}

	return res
}
