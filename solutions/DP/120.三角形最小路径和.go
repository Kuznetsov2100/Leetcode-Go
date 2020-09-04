/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	res := 1 << 31-1
	row := len(triangle)
	for i := 1; i < row; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
				continue
			}
			if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
				continue
			}
			triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
		}
	}
	for _, num := range triangle[row-1] {
		res = min(res, num)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

