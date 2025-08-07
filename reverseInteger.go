func reverse(x int) int {
	myStr := strconv.Itoa(x)
	isNeg := false
	if strings.ContainsRune(myStr, '-') {
		isNeg = true
		myStr = strings.ReplaceAll(myStr, "-", "")
	}
	runes := []rune(myStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	str := string(runes)
	fmt.Println(str)
	num, _ := strconv.Atoi(str)
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	if isNeg {
		return 0 - num
	} else {
		return num
	}
}