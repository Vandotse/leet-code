func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	start := intervals[0][0]
	end := intervals[0][1]
	var resp [][]int
	for pos, _ := range intervals {
		if pos < len(intervals)-1 {
			if end >= intervals[pos+1][0] {
				if end <= intervals[pos+1][1] {
					end = intervals[pos+1][1]
				}
			} else {
				resp = append(resp, []int{start, end})
				start = intervals[pos+1][0]
				end = intervals[pos+1][1]
			}
		} else {
			resp = append(resp, []int{start, end})
		}
	}
	return resp
}