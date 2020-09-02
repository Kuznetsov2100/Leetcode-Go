/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// grid[i][j]的含义是在位置(i，j)时的最小路径和
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 { // 第一行
				grid[i][j] = grid[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 { // 第一列
				grid[i][j] = grid[i-1][j] + grid[i][j]
				continue
			}
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]		
		}
	}
	return grid[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

