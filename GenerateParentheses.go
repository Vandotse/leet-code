func generateParenthesis(n int) []string {
	var resp []string
	start := "("
	var dfs func(input string, close int, open int)
	dfs = func(input string, close int, open int) {
		if close > open {
			return
		}
		if open > n {
			return
		}
		if open == close && open == n {
			resp = append(resp, input)
			return
		}
		dfs(input+")", close+1, open)
		dfs(input+"(", close, open+1)
	}

	dfs(start, 0, 1)
	return resp
}