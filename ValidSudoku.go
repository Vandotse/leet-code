func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	sqrs := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		sqrs[i] = make(map[byte]bool)
	}

	for row, curRow := range board {
		for col, _ := range curRow {
			val := board[col][row]

			if val == '.' {
				continue
			}

			if rows[row][val] {
				return false
			}
			rows[row][val] = true

			if cols[col][val] {
				return false
			}
			cols[col][val] = true

			square := (row/3)*3 + col/3
			if sqrs[square][val] {
				return false
			}
			sqrs[square][val] = true
		}
	}
	return true
}