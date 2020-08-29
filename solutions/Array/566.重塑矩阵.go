/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}
	row, col := len(nums), len(nums[0])
	if r * c != row * col {
		return nums
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	// 一维下标映射到二维下标
	for i := 0; i < row*col; i++ {
		res[i/c][i%c] = nums[i/col][i%col]
	}
	return res
}
// @lc code=end

