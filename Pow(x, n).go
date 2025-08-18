func myPow(x float64, n int) float64 {

	var helper func(x float64, n int) float64
	helper = func(x float64, n int) float64 {
		if n == 0 {
			return float64(1)
		}
		if n%2 == 0 {
			half := helper(x, n/2)
			return half * half
		}
		return x * helper(x, n-1)
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}
	resp := helper(x, n)

	return resp

}