package main

import (
	"testing"
)

type TestCase struct {
	data  string
	expected string
}

func testCases() []TestCase {
	return []TestCase{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"cbbdqwertyuio",
			"bb",
		},
	}
}

func TestLongestPalindrome(t *testing.T) {
	cases := testCases()

	for caseNum, item := range cases {
		res := longestPalindrome(item.data)
		if res != item.expected {
			t.Errorf("case: [%d] does not match, Got: %s. Expected: %s", caseNum, res, item.expected)
		}
	}
}
