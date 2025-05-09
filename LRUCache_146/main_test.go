package main

import (
	"reflect"
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	return reflect.DeepEqual(a, b)
}

type TestCase struct {
	instructions []string
	input        [][]int
	expected     []int
}

func testCases() []*TestCase {
	return []*TestCase{
		{
			instructions: []string{"LRUCache", "put", "get", "put", "get", "get"},
			input:        [][]int{[]int{1}, []int{2, 1}, []int{2}, []int{3, 2}, []int{2}, []int{3}},
			expected:     []int{1, -1, 2},
		},
		{
			instructions: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			input:        [][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}},
			expected:     []int{1, -1, -1, 3, 4},
		},
	}
}

func TestLRUCache(t *testing.T) {
	cases := testCases()

	for caseNum, testCase := range cases {
		var lruCache LRUCache
		var result []int
		for i, instruction := range testCase.instructions {
			switch instruction {
			case "LRUCache":
				lruCache = NewLRUCache(uint32(testCase.input[i][0]))
			case "put":
				lruCache.Put(testCase.input[i][0], testCase.input[i][1])
			case "get":
				result = append(result, lruCache.Get(testCase.input[i][0]))
			}
		}

		if !equal(result, testCase.expected) {
			t.Errorf("case: [%d] does not match, Got: %v. Expected: %v", caseNum, result, testCase.expected)
		}
	}
}
