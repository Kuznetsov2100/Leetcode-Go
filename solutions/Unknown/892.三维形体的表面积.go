/*
 * @lc app=leetcode.cn id=892 lang=golang
 *
 * [892] 三维形体的表面积
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	// 表面积=立方体个数×6−三种重叠部分的数量×2
	n, cubeCount, verticalOverlap, rowOverlap, colOverlap := len(grid), 0, 0, 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cubeCount += grid[i][j]
			if grid[i][j] > 1 {
				verticalOverlap += grid[i][j]-1
			}
			if i > 0 {
				rowOverlap += min(grid[i-1][j], grid[i][j])
			}
			if j > 0 {
				colOverlap += min(grid[i][j-1], grid[i][j])
			}
		}
	}
	return cubeCount*6- 2*(verticalOverlap+rowOverlap+colOverlap)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

