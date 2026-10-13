func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    
    for i := range len(nums) {
        val := nums[i]
        _, ok := m[val]

        if ok {
            return true
        }
        m[val] = struct{}{}
    }

    return false
}
