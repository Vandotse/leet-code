func combinationSum2(candidates []int, target int) [][]int {
	var resp [][]int
	sort.Ints(candidates)

	var dfs func(index int, current []int, total int)
	dfs = func(index int, current []int, total int) {

		if total == target {
			cpy := make([]int, len(current))
			copy(cpy, current)
			resp = append(resp, cpy)
			return
		}
		if index == len(candidates) || total > target {
			return
		}

		dfs(index+1, append(current, candidates[index]), total+candidates[index])

		for index+1 < len(candidates) && candidates[index] == candidates[index+1] {
			index += 1
		}
		dfs(index+1, current, total)
	}
	dfs(0, []int{}, 0)
	return resp
}
