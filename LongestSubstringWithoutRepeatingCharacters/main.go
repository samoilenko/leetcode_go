package main

import (
"fmt"
)

func main() {
	//s:= "ad日本語v日dfg"
	s:= "pwwkew"
	//maxLength:=0
	//
	//current := make(map[rune]int)
	//
	//for i, c:= range s {
	//	if curIndex, ok := current[c]; ok {
	//		if len(current) > maxLength {
	//			maxLength = len(current)
	//		}
	//		current = make(map[rune]int)
	//		for ni, k := range s[curIndex: i] {
	//			current[k] = ni
	//		}
	//		//for j := curIndex; j < i; {
	//		//
	//		//}
	//	} else {
	//		current[c] = i
	//	}
	//}

	fmt.Println(lengthOfLongestSubstring(s))
}


func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	current := make(map[rune]int)


	for i, c:= range s {
		if curIndex, ok := current[c]; ok {
			if len(current) > maxLength {
				maxLength = len(current)
			}
			current = make(map[rune]int)
			for ni, k := range s[curIndex: i] {
				current[k] = ni
			}
			//for j := curIndex; j < i; {
			//
			//}
		} else {
			current[c] = i
		}
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