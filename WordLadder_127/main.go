package main

func hasOneLetterDifference(word1, word2 string) bool {
	var count int
	for i := range word1 {
		if word1[i] != word2[i] {
			count += 1
			if count > 1 {
				return false
			}
		}
	}

	return count == 1
}

func createGraph(list []string) map[int][]int {
	graph := make(map[int][]int)
	for i, word1 := range list {
		graph[i] = []int{}
		for j, word2 := range list {
			if hasOneLetterDifference(word1, word2) {
				graph[i] = append(graph[i], j)
			}
		}
	}

	return graph
}

// uses BFS to find the shortest path between two words
// https://leetcode.com/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	endWordIndex := -1
	for i, word := range wordList {
		if word == endWord {
			endWordIndex = i
			break
		}
	}
	if endWordIndex == -1 {
		return 0
	}
	graph := createGraph(wordList)
	l := len(wordList)
	visited := make([]bool, l)
	visited[l-1] = true
	queue := make([]int, 1)
	queue[0] = l - 1
	distance := 0

	for len(queue) > 0 {
		childrenCount := len(queue)
		distance++
		for i := 0; i < childrenCount; i++ {
			wordIndex := queue[0]
			queue = queue[1:]

			if endWordIndex == wordIndex {
				return distance
			}

			for _, nextWordIndex := range graph[wordIndex] {
				if !visited[nextWordIndex] {
					visited[nextWordIndex] = true
					queue = append(queue, nextWordIndex)
				}
			}
		}
	}

	return 0
}
