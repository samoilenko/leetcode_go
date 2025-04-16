package main

import (
	"reflect"
	"sort"
	"testing"
)

func equalUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	// Make copies to avoid mutating original slices
	aCopy := append([]int(nil), a...)
	bCopy := append([]int(nil), b...)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	return reflect.DeepEqual(aCopy, bCopy)
}

type TestCase struct {
	data     []int
	k        int
	expected []int
}

func testCases() []TestCase {
	return []TestCase{
		{
			[]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3},
			3,
			[]int{3, 7, 10},
		},
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
	}
}

func TestTopKFrequentElements(t *testing.T) {
	cases := testCases()

	for caseNum, item := range cases {
		res := topKFrequent(item.data, item.k)
		if !equalUnordered(res, item.expected) {
			t.Errorf("case: [%d] does not match, Got: %v. Expected: %v", caseNum, res, item.expected)
		}
	}
}
