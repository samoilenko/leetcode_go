package main

import (
	"fmt"
)

func main() {
	//s:= "ad日本語v日dfg"
	//s:= "yfsrsrpzuya"
	//s:= "dvdf"
	s:= "aabaab!bb"

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	current := make(map[rune]int)
	lastIndex := 0


	for i, c:= range s {
		if curIndex, ok := current[c]; ok {
			if len(current) > maxLength {
				maxLength = len(current)
			}
			//current = make(map[rune]int)

			//for _, k := range s[curIndex: i] {
			for delIndex, k := range s[lastIndex:curIndex] {
				if _, ok := current[k]; ok {
					if delIndex >= current[k] {
						delete(current, k)
					}
				}
				//current[k] = ni
			}

			lastIndex = curIndex

			//delete(current, c)
			//for j := curIndex; j < i; {
			//
			//}
		}

		current[c] = i
	}

	if len(current) > maxLength {
		return len(current)
	}

	return maxLength
}

//func remove subsequence(s string, index int) {
//for i, c:= range s {

//}
//}