/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	transform := make(map[string][]int)
	for index, word := range wordList {
		for i := range word {
			k := word[:i] + "*" + word[i+1:]
			if _, ok := transform[k]; !ok {
				transform[k] = []int{}
			}
			transform[k] = append(transform[k], index)
		}
	}

	count := 1
	visited := make([]bool, len(wordList))
	queue := []string{beginWord}
	
	for len(queue) > 0 {
		count++
		var newQueue[]string
		for _, word := range queue {
			for i := range word {
				key := word[:i] + "*" + word[i+1:]
				for _, index := range transform[key] {
					if  wordList[index] == endWord {
						return count
					}
					if !visited[index] {
						visited[index] = true
						newQueue = append(newQueue, wordList[index])
					}
				}
			}
		}
		queue = newQueue
	}
	return 0
}
// @lc code=end

