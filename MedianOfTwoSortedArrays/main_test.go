package main

import (
	"testing"
)

type TestCase struct {
	nums1 []int
	nums2 []int
	expected float64
}

func testCases()[]*TestCase {
	return []*TestCase{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			expected: 2,
		},
		{
			nums1: []int{1, 4},
			nums2: []int{2, 3},
			expected: 2.5,
		},
		{
			nums1: []int{2},
			nums2: []int{},
			expected: 2,
		},
		{
			nums1: []int{1},
			nums2: []int{1},
			expected: 1,
		},
		{
			nums1: []int{},
			nums2: []int{2, 3},
			expected: 2.5,
		},
	}
}

func TestMedianOfTwoNumbers(t *testing.T) {
	cases := testCases()

	for caseNum, item := range cases {
		result := findMedianSortedArrays(item.nums1, item.nums2)
		if result != item.expected {
			t.Errorf("case: [%d] does not match, Got: %f. Expected: %f", caseNum, result, item.expected)
		}
	}
}
