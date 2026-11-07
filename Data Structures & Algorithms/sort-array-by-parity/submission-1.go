func sortArrayByParity(nums []int) []int {
	res:= []int{}

	even := []int{}
	odd := []int{}

	for _,num := range nums {
		if num % 2 == 0 {
			even = append(even,num)
		} else {
			odd = append(odd,num)
		}
	}


	for _, num := range even {
		res = append(res, num)
	}

	for _, num := range odd {
		res = append(res, num)
	}

	return res
}
