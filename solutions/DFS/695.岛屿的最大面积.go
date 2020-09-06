/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	var count int // 记录一个岛屿的面积
	var res int

	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 坐标（x,y）非法 或者是水域或者是已访问过的点，则返回
		if (x < 0 || x >= m || y < 0 || y >= n) || grid[x][y] != 1 {
			return
		}
		grid[x][y] = 2 // 已访问过的点标记为2
		count++ // 岛屿面积+1
		dfs(x+1, y)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x, y-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 { // 该点是陆地，且没有被标记过
				count  = 0 // 在计算一块新的岛屿面积之前，将count置零
				dfs(i, j)
				if count > res {
					res = count
				}
			}
		}
	}
	return res
}
// @lc code=end

