func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	mid := len(merged) / 2
	if len(merged)%2 == 0 {
		return (float64(merged[mid-1]) + float64(merged[mid])) / 2.0
	} else {
		return float64(merged[mid])
	}

}