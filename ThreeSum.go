package main

import "slices"

func threeSum(nums []int) [][]int {
	var resp [][]int
	slices.Sort(nums)
	for pos, value := range nums {
		if pos > 0 && value == nums[pos-1] {
			continue
		}
		l := pos + 1
		r := len(nums) - 1
		for l < r {
			threeSum := nums[r] + nums[l] + value
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				resp = append(resp, []int{value, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return resp
}
