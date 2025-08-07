func isMatch(s string, p string) bool {
	cashe := make(map[[2]int]bool)
	return dfs(s, p, 0, 0, cashe)
}

func dfs(s string, p string, i int, j int, cashe map[[2]int]bool) bool {

	b := [2]int{i, j}
	val, ok := cashe[b]
	if ok {
		return val
	}
	if i >= len(s) && j >= len(p) {
		return true
	}
	if j >= len(p) {
		return false
	}

	match := i < len(s) && (s[i] == p[j] || p[j] == '.')

	if j+1 < len(p) && p[j+1] == '*' {
		cashe[b] = (match && dfs(s, p, i+1, j, cashe)) || dfs(s, p, i, j+2, cashe)
		return cashe[b]
	}
	if match {
		cashe[b] = dfs(s, p, i+1, j+1, cashe)
		return cashe[b]
	}
	cashe[b] = false
	return cashe[b]
}