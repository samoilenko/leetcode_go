package main

import (
	"fmt"
)

func main() {
	s := "aabaab!bb" // 3
	fmt.Println(withoutDeleting(s))
}

func withoutDeleting(s string) int {
	var maxLength int = 0
	current := make(map[rune]int)
	j := 0
	lastIndex := 0
	for i, c := range s {
		if curIndex, ok := current[c]; ok && curIndex >= lastIndex {
			j = i - curIndex
			lastIndex = curIndex
		} else {
			j += 1
		}

		if j > maxLength {
			maxLength = j
		}

		current[c] = i
	}

	return maxLength
}

func lengthOfLongestSubstringWindowFrame(s string) int {
	var maxLength int = 0
	current := make(map[rune]int)
	lastIndex := 0
	for i, c := range s {
		if curIndex, ok := current[c]; ok {
			if maxLength < len(current) {
				maxLength = len(current)
			}
			for _, v := range s[lastIndex:curIndex] {
				delete(current, v)
			}

			lastIndex = curIndex + 1
		}

		current[c] = i
	}

	if maxLength < len(current) {
		maxLength = len(current)
	}

	return maxLength
}

func lengthOfLongestSubstring(s string) int {
	var maxLength int = 0
	current := make(map[rune]int)

	for i, c := range s {
		if curIndex, ok := current[c]; ok {
			if len(current) > maxLength {
				maxLength = len(current)
			}

			removeExtraChars(current, curIndex)
		}

		current[c] = i
	}

	if len(current) > maxLength {
		return len(current)
	}

	return maxLength
}

func removeExtraChars(chars map[rune]int, index int) {
	for i, v := range chars {
		if index >= v {
			delete(chars, i)
		}
	}
}

//func remove subsequence(s string, index int) {
//for i, c:= range s {

//}
//}
