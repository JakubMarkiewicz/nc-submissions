func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[r] < prices[l] {
			l = r
		} else {
			if tmpMaxP := prices[r] - prices[l]; tmpMaxP > maxP {
				maxP = tmpMaxP
			}
		}
		r++
	}

	return maxP
}
