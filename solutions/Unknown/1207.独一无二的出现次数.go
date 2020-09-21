/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	m1 := make([]int, 2001)
	for i := range arr {
		m1[arr[i]+1000]++
	}
	
	m2 := make([]int, 1001)
	for i := range m1 {
		if m1[i] > 0 {
			m2[m1[i]]++
		}
	}

	for i := range m2 {
		if m2[i] > 1 {
			return false
		}
	}
	return true
}
// @lc code=end

