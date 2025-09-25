func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if row[len(row)-1] >= target {
			for _, val := range row {
				if val == target {
					return true
				}
			}
			return false
		}
	}
	return false
}