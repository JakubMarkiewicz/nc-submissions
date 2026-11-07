import (
	"slices"
)

func sortArrayByParity(nums []int) []int {
	slices.SortFunc(nums, func(a,b int) int {
		if (a % 2 == 0 && b % 2 == 0) || (a % 2 != 0 && b % 2 != 0 ) {
			if a > b {
				return 1
			} 
				return -1
		}  else if a % 2 == 0 {
			return -1
		}
			return 1
	})	
	return nums
}
