func change(amount int, coins []int) int {
	resp := make([]int, amount+1)
	resp[0] = 1
	for _, value := range coins {
		for pos, _ := range resp {
			if pos-value >= 0 {
				resp[pos] = resp[pos] + resp[pos-value]
			}
		}
	}
	return resp[len(resp)-1]
}