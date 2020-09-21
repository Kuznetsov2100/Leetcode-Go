/*
 * @lc app=leetcode.cn id=1170 lang=golang
 *
 * [1170] 比较字符串最小字母出现频次
 */

// @lc code=start
func numSmallerByFrequency(queries []string, words []string) []int {
	// 透过现象看本质：存在数组A，B，遍历数组A, 对于A[i],找到数组B中大于A[i]的数字的个数
	m, n := len(queries), len(words)
	queriesFreq := make([]int, m)
	for i := range queries {
		queriesFreq[i] = f(queries[i])
	}

	wordsFreq := make([]int, n)
	for i := range words {
		wordsFreq[i] = f(words[i])
	}
	sort.Ints(wordsFreq)

	res := make([]int, m)
	for i := range queriesFreq {
		index := sort.Search(n, func(j int) bool {return wordsFreq[j] > queriesFreq[i]})
		if index < n {
			res[i] = n-index
		}
	}
	return res
}

func f(s string) int {
	n := len(s)
	count := 1
	minchar := s[0]
	for i := 1; i < n; i++ {
		if s[i] == minchar {
			count++
		} else if s[i] < minchar {
			minchar = s[i]
			count = 1
		}
	}
	return count
}
// @lc code=end

