func getConcatenation(nums []int) []int {
   res := make([]int, len(nums) * 2) 


   for i := 0; i < len(nums); i++ {
		res[i] = nums[i]
		res[i+len(nums)] = nums[i]
   }

   return res
}
