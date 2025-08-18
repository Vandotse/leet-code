func rotate(matrix [][]int) {
	var stack [][]int
	for _, val := range matrix {
		cpy := make([]int, len(val))
		copy(cpy, val)
		stack = append(stack, cpy)
	}
	for i, _ := range stack {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for j, _ := range value {
			matrix[j][i] = value[j]
		}
	}
}