package main

import (
    "reflect"
    "runtime"
    "strings"
    "testing"
)

type TestCase struct {
    data  string
    count int
}

func testCases() []TestCase {
    return []TestCase{
        {"dvdf", 3},
        {"yfsrsrpzuya", 7},
        {"aabaab!bb", 3},
        {"jbpnbwwd", 4},
        {"bpfbhmipx", 7},
        {"abcabcbb", 3},
    }
}

func testFunctions() []func(s string) int {
    return []func(s string) int{
        withoutDeleting,
        lengthOfLongestSubstringWindowFrame,
        lengthOfLongestSubstring,
    }
}

func TestAllAvailableFunctions(t *testing.T) {
    cases := testCases()

    for _, currentFunc := range testFunctions() {
        funcName := runtime.FuncForPC(reflect.ValueOf(currentFunc).Pointer()).Name()

        for caseNum, item := range cases {
            countLetters := currentFunc(item.data)
            if countLetters != item.count {
                t.Errorf("Function %s, case: [%d] does not match, Got: %d. Expected: %d",
                    funcName[strings.LastIndexAny(funcName, ".")+1:], caseNum, countLetters, item.count)
            }
        }
    }
}
