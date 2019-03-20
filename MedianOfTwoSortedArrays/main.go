package main

import (
    "fmt"
    "math"
)

func main()  {
    a := findMedianSortedArrays([]int{}, []int{2, 3})

    fmt.Println(a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n := len(nums1)
    m := len(nums2)

    if n == 0 && m == 0 {
        return 0
    }

    if n == 0 {
        return findCentral(nums2)
    }
    if m == 0 {
        return findCentral(nums1)
    }

    l := n + m
    i, j := 0, 0
    var median, oldMedian float64

    for l >= 0 {
        oldMedian = median
        if i < n && (j >= m || nums1[i] < nums2[j]) {
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

func findCentral(nums []int) float64 {
    l := len(nums)
    if l % 2 == 0 {
        var averageIndex int = l / 2
        return float64( float64(nums[averageIndex] + nums[averageIndex - 1]) / 2 )
    }

    centralIndex := math.Floor(float64(len(nums) / 2))

    return float64(nums[int(centralIndex)])
}
