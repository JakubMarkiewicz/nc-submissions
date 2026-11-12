type KthLargest struct {
	stack []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)

	stack := nums
	if len(stack) > k {
		stack = stack[len(nums)-k:]
	}

	return KthLargest{
		stack,
		k,
	}	
}


func (this *KthLargest) Add(val int) int {
	if len(this.stack) < this.k {
		this.stack = append(this.stack, val)
		sort.Ints(this.stack)
	} else if val > this.stack[0] {
		idx := 0;
		for idx < len(this.stack) {
			if val < this.stack[idx] {
				break
			}
			idx++
		}
		rest := append([]int(nil), this.stack[idx:]...)
		this.stack = append(this.stack[1:idx], val)
		this.stack = append(this.stack, rest...)
	}

	return this.stack[0]
}