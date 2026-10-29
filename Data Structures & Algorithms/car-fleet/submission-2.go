type posIdx struct {
	p int
	s int
}

func carFleet(target int, position []int, speed []int) int {
	pos := []posIdx{}

	for i, value := range position {
		pos = append(pos, posIdx{s: speed[i], p: value})
	}

	sort.Slice(pos, func (i,j int) bool {
		return pos[i].p > pos[j].p
	})

	movesStack := []float64{}

	for _, value := range pos {
		movesToTarget := float64(target - value.p) / float64(value.s)


	if len(movesStack) == 0 ||
		movesToTarget > movesStack[len(movesStack)-1] {
		movesStack = append(movesStack, movesToTarget)
	}
		
	}


	return len(movesStack)
}
