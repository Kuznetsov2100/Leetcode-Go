/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	greatest := -1
	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] > greatest {
			arr[i], greatest = greatest, arr[i]
		} else {
			arr[i] = greatest
		}
	}
	return arr
}
// @lc code=end

