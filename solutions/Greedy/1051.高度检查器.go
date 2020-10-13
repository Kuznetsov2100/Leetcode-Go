/*
 * @lc app=leetcode.cn id=1051 lang=golang
 *
 * [1051] 高度检查器
 */

// @lc code=start
func heightChecker(heights []int) int {
	var res int
	arr := make([]int, len(heights))
	copy(arr, heights)
	sort.Ints(arr)
	for i, height := range heights {
		if height != arr[i] {
			res++
		}
	}
	return res
}
// @lc code=end

