func topKFrequent(nums []int, k int) []int {
 fMap := make(map[int]int)
 freq := make([][]int, len(nums)+1)

 for  _, n := range nums {
	fMap[n] += 1
 }

 for k, v := range fMap {
	freq[v] = append(freq[v], k)
 }

 res := []int{}

 for i := len(freq) - 1; i >= 0; i-- {
	if len(freq[i]) != 0 {
		res = append(res, freq[i]...)
	}

	if len(res) == k {
		break
	}
 }

 return res[0:k]
}
