func isValid(s string) bool {
	var stack []rune
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, val := range s {
		match, found := hash[val]
		if found {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, val)
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}