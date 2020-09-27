/*
 * @lc app=leetcode.cn id=1394 lang=golang
 *
 * [1394] 找出数组中的幸运数
 */

// @lc code=start
func findLucky(arr []int) int {
	res := -1
	m := make([]int, 501)
	for i := range arr {
		m[arr[i]]++
	}

	for i, count := range m {
		if i == 0 || count == 0 {
			continue
		}
		if i == count {
			res = i
		}
	}
	return res
}
// @lc code=end

