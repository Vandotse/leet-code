package main

import "slices"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	perm := permute(nums[:1])
	var resp [][]int
	for p := range perm {
		for i := 0; i < len(p)+1; i++ {
			cpy := make([]int, len(p)+1)
			copy(cpy, p)
			cpy = slices.Insert(cpy, i, nums[0])
			resp = append(resp, cpy)
		}
	}
	return resp
}
