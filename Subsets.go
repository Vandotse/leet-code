func subsets(nums []int) [][]int {
	var resp [][]int
	var dfs func(sub []int, pos int)
	dfs = func(sub []int, pos int) {
		if pos == len(nums) {
			copySub := append([]int{}, sub...)
			resp = append(resp, copySub)
			return
		}
		dfs(sub, pos+1)
		dfs(append(sub, nums[pos]), pos+1)
	}
	dfs([]int{}, 0)
	return resp
}