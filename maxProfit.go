package main

// solved
func maxProfit(prices []int) int {
	maxProfit := 0
	buy := prices[0]
	for _, sell := range prices {
		profit := sell - buy

		if sell < buy {
			buy = sell
		} else if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
