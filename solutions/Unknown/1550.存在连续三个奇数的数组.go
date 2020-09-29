/*
 * @lc app=leetcode.cn id=1550 lang=golang
 *
 * [1550] 存在连续三个奇数的数组
 */

// @lc code=start
func threeConsecutiveOdds(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	odds := 0
	for i := 0; i < n; i++ {
		if arr[i] % 2 == 0 {
			odds = 0
			continue
		}
		odds++
		if odds == 3 {
			return true
		}
	}
	return false
}
// @lc code=end

