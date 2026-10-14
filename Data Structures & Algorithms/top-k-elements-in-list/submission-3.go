func insert(arr []int, idx int, val int) []int {
	arr = append(arr, 0)
	copy(arr[idx + 1:], arr[idx:])
	arr[idx] = val
	return arr
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num] += 1
	}

	sorted := []int{}

	for key, v := range m {
		if len(sorted) == 0 {
			sorted = append(sorted, key)
			continue
		}


		match := false
		for i, sV := range sorted {
			if v >= m[sV] {
				sorted = insert(sorted,i,key)
				match = true
				break
			}
		}		

		if !match && len(sorted) < key {
			sorted = append(sorted, key)
		}
	}



	return sorted[0:k]

}
