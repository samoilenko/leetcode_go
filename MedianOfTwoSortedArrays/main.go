package main

func main()  {
	
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	l := n + m
	i, j := 0, 0
	var median, oldMedian float64

	for l > 0 {
		oldMedian = median
		if i < n && nums1[i] < nums2[j] {
			median = float64(nums1[i])
			i += 1
		} else {
			median = float64(nums2[j])
			j += 1
		}

		l -= 2
	}

	if l % 2 == 0 {
		median = (median + oldMedian) / 2
	}

	return median
}
