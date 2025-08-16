func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n, m := len(num1), len(num2)
	result := make([]int, n+m)

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			digit1 := int(num1[i] - '0')
			digit2 := int(num2[j] - '0')
			multi := digit1 * digit2
			place1, place2 := i+j, i+j+1

			result[place2] += multi
			result[place1] += result[place2] / 10
			result[place2] = result[place2] % 10
		}
	}

	zeros := 0
	if result[zeros] == 0 {
		zeros++
	}
	var res strings.Builder
	for i := zeros; i < len(result); i++ {
		res.WriteByte(byte(result[i] + '0'))
	}

	return res.String()
}