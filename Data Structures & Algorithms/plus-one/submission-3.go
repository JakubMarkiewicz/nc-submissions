func plusOne(digits []int) []int {
   res := []int{} 

	over := false
   for idx := len(digits) - 1; idx >= 0; idx-- {
	value := digits[idx]
	if over || len(res) == 0 {
		value += 1
	}
	over = value == 10
	if len(res) == 0 {
		res = prepend(res, value % 10)
		continue
	}
	if res[0] == 0 {
		res = prepend(res, value % 10)
	} else {
		res = prepend(res, value)
	}
   }


   if over {
	res = prepend(res, 1)
   }

   return res
}


func prepend(arr []int, v int) []int  {
	return append([]int{v}, arr...)
}