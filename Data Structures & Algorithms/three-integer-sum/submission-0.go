import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	res := [][]int{}

	for idx, num := range nums {
		m[num]--

		if idx > 0 && num == nums[idx-1] {
			continue
		}

		for j := idx + 1; j < len(nums); j++ {
			m[nums[j]]--

			if j > idx + 1 && nums[j] == nums[j-1] {
				continue
			}
			needed := -(num + nums[j])

			if m[needed] > 0 {
				res = append(res, []int{num, nums[j], needed})	
			}
		}

		for j := idx + 1; j < len(nums); j++ {
			m[nums[j]]++
		}
	}

	return res

}
