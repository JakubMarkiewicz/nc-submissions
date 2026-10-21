func sortColors(nums []int) {
   m := map[int]int{0: 0, 1: 0, 2:0}

  for _, num := range nums {
	m[num] += 1
  }

  x := 0
  for i := 0; i < len(nums); i++ {
	if m[x] == 0 {
		x++
		i--
		continue
	}

		nums[i] = x
		m[x] -= 1
  }
  
}
