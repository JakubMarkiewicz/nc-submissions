func maxProfit(prices []int) int {
	profit := 0
	lowest := 0

	for idx, price := range prices {
		if price < prices[lowest] {
			lowest = idx	
			continue
		}

		potentialProfit := price - prices[lowest]
		if potentialProfit > profit {
			profit = potentialProfit
		}

	}

	return profit
}
