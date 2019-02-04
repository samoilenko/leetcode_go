package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	var res []int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i+1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return append(res, i, j)
// 			}
// 		}

// 	}

// 	return res
// }

func twoSum(nums []int, target int) []int {
	copy := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		nmb := target - nums[i]

		if _, ok := copy[nmb]; ok {
			return []int{copy[nmb], i}
		}

		copy[nums[i]] = i
	}

	return []int{}
}

func main() {
	fmt.Prinyln(twoSum([]int{2, 7, 11, 15}, 9))
}
