func groupAnagrams(strs []string) [][]string {
	resp := make(map[string][]string)
	var res [][]string
	for _, str := range strs {
		runes := []rune(str)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedStr := string(runes)
		resp[sortedStr] = append(resp[sortedStr], str)
	}

	for _, val := range resp {
		res = append(res, val)
	}

	return res
}