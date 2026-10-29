func searchMatrix(matrix [][]int, target int) bool {
	l,r := 0, len(matrix) * len(matrix[0]) - 1

	for l < r {
		middle := l + (r - l) / 2
		row := middle / len(matrix[0])
		col := middle % len(matrix[0])

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			r = middle
		} else {
			l = middle + 1
		}

	}



		return matrix[l/len(matrix[0])][l%len(matrix[0])] == target

}
