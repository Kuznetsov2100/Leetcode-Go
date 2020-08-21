/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var res [][]int
	for i := 0; i <= rowIndex; i++ {
		arr := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				arr[j] = 1
			} else {
				arr[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, arr)
	}
	return res[rowIndex]
}
// @lc code=end

