func majorityElement(nums []int) int {
   m := make(map[int]int) 

   for _, num := range nums {
	 m[num]++

	 if m[num] > (len(nums) / 2) {
		return num
	 } 
   }

   return 0
}
