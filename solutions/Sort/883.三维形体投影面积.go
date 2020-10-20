/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] 三维形体投影面积
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	// 问题转化：总投影面积 = xy+yz+xz
	// xy平面的投影面积为 grid[i][j] > 0 的个数
	// yz平面的投影面积为每一行的最大值的和
	// xz平面的投影面积为每一列的最大值的和
	n, xy, yz, xz := len(grid), 0, 0, 0
	for i := 0; i < n; i++ {
		rowMax := grid[i][0]
		colMax := grid[0][i]
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				xy++		
			}
			if j > 0 {
				rowMax = max(rowMax, grid[i][j])
				colMax = max(colMax, grid[j][i])
			}	
		}
		yz += rowMax
		xz += colMax
	}
	return xy+yz+xz
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

