func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	stack := []int{}

	for idx, temp := range temperatures {

		for len(stack) > 0 && temp > temperatures[stack[len(stack) - 1]] {
			last := stack[len(stack) - 1]
			res[last] = idx - last
			stack = stack[0: len(stack) - 1]

		}


		stack = append(stack,idx)
	}

	return res
}
