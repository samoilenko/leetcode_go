package main

import (
	"testing"
)

type TestCase struct {
	data  string
	count int
}

func testCases() []TestCase {
	return []TestCase{

	}
}

func TestAllAvailableFunctions(t *testing.T) {
	cases := testCases()

	for caseNum, item := range cases {
		countLetters := currentFunc(item.data)
		if countLetters != item.count {
			t.Errorf("Function %s, case: [%d] does not match, Got: %d. Expected: %d",
				funcName[strings.LastIndexAny(funcName, ".")+1:], caseNum, countLetters, item.count)
		}
	}
}
