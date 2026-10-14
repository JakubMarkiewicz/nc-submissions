func getConcatenation(nums []int) []int {
   res := make([]int, len(nums) * 2) 

   for i := 0; i < len(res); i++ {
		res[i] = nums[i % len(nums)]
   }

   return res
}
