package main

func jump(nums []int) int {
	resp := 0
	l, r := 0, 0
	for r < len(nums)-1 {
		farthest := 0
		for i := l; i <= r; i++ {
			farthest = max(farthest, i+nums[i])
		}
		l = r + 1
		r = farthest
		resp += 1
	}
	return resp
}
