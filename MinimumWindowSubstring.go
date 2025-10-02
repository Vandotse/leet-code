func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}
	substr := make(map[rune]int)
	window := make(map[rune]int)
	for _, val := range t {
		_, exists := substr[val]
		if exists {
			substr[val] += 1
		} else {
			substr[val] = 1
		}
	}
	have := 0
	need := len(substr)
	res := []int{-1, -1}
	reslen := math.MaxInt64
	l := 0
	for r, val := range s {
		_, exists := window[val]
		if exists {
			window[val] += 1
		} else {
			window[val] = 1
		}
		_, isIn := substr[val]
		if isIn && window[val] == substr[val] {
			have++
		}

		for have == need {
			if r-l+1 < reslen {
				res = []int{l, r}
				reslen = r - l + 1
			}
			window[rune(s[l])] -= 1
			_, isIn := substr[rune(s[l])]
			if isIn && window[rune(s[l])] < substr[rune(s[l])] {
				have--
			}
			l++
		}
	}
	if reslen == math.MaxInt64 {
		return ""
	}
	ans := s[res[0] : res[1]+1]
	return ans
}