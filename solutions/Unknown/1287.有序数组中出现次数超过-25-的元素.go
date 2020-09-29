/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	var res int
	threshold := len(arr)/4
	for i := range arr {
		if arr[i] == arr[i+threshold] {
			res = arr[i]
			break
		}
	}
	return res
}
// @lc code=end

