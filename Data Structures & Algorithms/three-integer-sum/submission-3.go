import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	fmt.Printf("%d", nums)

	res := [][]int{}

	count := make(map[int]int)

	for _, value := range nums {
		count[value]++
	}

	for idx := 0; idx < len(nums); idx++ {
		count[nums[idx]]--

		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		for jIdx := idx + 1; jIdx < len(nums); jIdx++ {
			count[nums[jIdx]]--

			if jIdx > idx + 1 && nums[jIdx] == nums[jIdx - 1] {
				continue
			}

			needed := (nums[idx] + nums[jIdx]) * -1

			if count[needed] > 0 {
				res = append(res,[]int{nums[idx], nums[jIdx], needed})
			}

		}

		for jIdx := idx + 1; jIdx < len(nums); jIdx++ {
			count[nums[jIdx]]++
		}
	}
	return res
}
