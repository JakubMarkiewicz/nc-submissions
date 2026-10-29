func carFleet(target int, position []int, speed []int) int {
	pos := [][]int{}

	for i, value := range position {
		pos = append(pos, []int{speed[i], value})
	}

	sort.Slice(pos, func (i,j int) bool {
		return pos[i][1] > pos[j][1]
	})

	movesStack := []float64{}

	for _, value := range pos {
		movesToTarget := float64(target - value[1]) / float64(value[0])


	if len(movesStack) == 0 ||
		movesToTarget > movesStack[len(movesStack)-1] {
		movesStack = append(movesStack, movesToTarget)
	}
		
	}


	return len(movesStack)
}
