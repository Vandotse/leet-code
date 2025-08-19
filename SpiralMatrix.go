func spiralOrder(matrix [][]int) []int {
	var resp []int
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			resp = append(resp, matrix[top][col])
		}
		top++
		for row := top; row <= bottom; row++ {
			resp = append(resp, matrix[row][right])
		}
		right--
		if top <= bottom {
			for col := right; col >= left; col-- {
				resp = append(resp, matrix[bottom][col])
			}
			bottom--
		}
		if left <= right {
			for row := bottom; row >= top; row-- {
				resp = append(resp, matrix[row][left])
			}
			left++
		}
	}
	return resp
}