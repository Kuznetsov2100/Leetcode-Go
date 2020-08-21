/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var pre []int
	for i := 0; i <= rowIndex; i++ {
		arr := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				arr[j] = 1
			} else {
				arr[j] = pre[j-1] + pre[j]
			}
		}
		pre = make([]int, i+1)
		copy(pre, arr)
	}
	return pre
}
// @lc code=end

