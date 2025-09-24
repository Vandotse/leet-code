func setZeroes(matrix [][]int) {
	var rows []int
	var cols []int
	for rowIdx, row := range matrix {
		for colIdx, val := range row {
			if val == 0 {
				rows = append(rows, rowIdx)
				cols = append(cols, colIdx)
			}
		}
	}
	for _, zeroRow := range rows {
		for pos, _ := range matrix[zeroRow] {
			matrix[zeroRow][pos] = 0
		}
	}
	for _, zeroCol := range cols {
		for x := range matrix {
			matrix[x][zeroCol] = 0
		}
	}
}