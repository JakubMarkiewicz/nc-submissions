
type NumMatrix struct {
	value [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{value: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	r,c := row1, col1

	for r <= row2 && c <= col2 {
		sum += this.value[r][c]

		if c == col2 {
			r++
			c = col1
		} else {
			c++
		}
	} 

	return sum

}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
