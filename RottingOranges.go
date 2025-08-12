func orangesRotting(grid [][]int) int {
	var queue [][2]int
	time := 0
	fresh := 0
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) != 0 && fresh > 0 {
		directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for _, rotten := range queue {
			for _, dir := range directions {
				row := rotten[0] + dir[0]
				col := rotten[1] + dir[1]
				if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || grid[row][col] != 1 {
					continue
				}
				grid[row][col] = 2
				queue = append(queue, [2]int{row, col})
				fresh -= 1
			}
			queue = queue[1:]
		}
		time++
	}
	if fresh > 0 {
		return -1
	} else {
		return time
	}
}