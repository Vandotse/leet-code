func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var resp []string
	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtrack func(index int, curStr string)
	backtrack = func(index int, curStr string) {
		if len(curStr) == len(digits) {
			resp = append(resp, curStr)
			return
		}
		str := phoneMap[digits[index]-'2']
		for _, value := range str {
			backtrack(index+1, curStr+string(value))
		}
	}
	backtrack(0, "")
	return resp
}
