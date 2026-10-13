func twoSum(nums []int, target int) []int {
   m := make(map[int]int)  

   for i, n := range nums {
	mVal, ok := m[target - n]

	if ok {
		return []int{mVal, i}
	}

	m[n] = i
   }

   return []int{}
}
