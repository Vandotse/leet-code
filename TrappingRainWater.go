func trap(height []int) int {
	resp := 0
	highestRight := make([]int, len(height))
	highestLeft := make([]int, len(height))
	maxL := 0
	maxR := 0
	for pos, _ := range height {
		if pos == 0 {
			highestLeft[pos] = 0
			continue
		}
		if height[pos-1] > maxL {
			maxL = height[pos-1]
		}
		highestLeft[pos] = maxL
	}
	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			highestRight[i] = 0
			continue
		}
		if height[i+1] > maxR {
			maxR = height[i+1]
		}
		highestRight[i] = maxR
	}
	for i := 0; i < len(height); i++ {
		minHighest := min(highestRight[i], highestLeft[i])
		water := minHighest - height[i]
		if water > 0 {
			resp = resp + water
		}
	}
	return resp
}