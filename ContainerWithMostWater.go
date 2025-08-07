func maxArea(height []int) int {
	largest := 0
	start := 0
	end := len(height) - 1
	for start < len(height) && end > 0 {
		if height[start] > height[end] {
			if largest < height[end]*(end-start) {
				largest = height[end] * (end - start)
			}
			end--
		} else {
			if largest < height[start]*(end-start) {
				largest = height[start] * (end - start)
			}
			start++
		}
	}
	return largest
}