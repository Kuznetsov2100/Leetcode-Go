/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	m := make([]int, 2001)
	for i := range arr {
		m[arr[i]]++
	}
	j := 1
	for ; j < 2001; j++ {
		if m[j] == 0 {
			k--
		}
		if k == 0 {
			break
		}
	}
	return j
}
// @lc code=end

