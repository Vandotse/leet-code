func maxSubArray(nums []int) int {
	total := 0
	max := nums[0]
	for _, val := range nums {
		if total < 0 {
			total = 0
		}
		total = total + val
		if total > max {
			max = total
		}
	}
	return max
}