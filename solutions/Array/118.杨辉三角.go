/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
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
	return res
}
// @lc code=end

