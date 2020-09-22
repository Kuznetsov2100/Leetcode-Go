/*
 * @lc app=leetcode.cn id=1254 lang=golang
 *
 * [1254] 统计封闭岛屿的数目
 */

// @lc code=start
func closedIsland(grid [][]int) int {
	var res int
	var countIsland func(x, y int)
	var floodFill func(x, y int)

	m, n := len(grid), len(grid[0])

	inArea := func(x, y int) bool {
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		return true
	}

	floodFill = func(x, y int) {
		if !inArea(x, y) || grid[x][y] != 0 {
			return
		}
		grid[x][y] = 1
		floodFill(x+1, y)
		floodFill(x, y+1)
		floodFill(x-1, y)
		floodFill(x, y-1)
	}

	countIsland = func(x, y int) {
		// 坐标（x,y）非法 或者是水域或者是已访问过的点，则返回
		if !inArea(x, y) || grid[x][y] != 0 {
			return
		}
		grid[x][y] = 2 // 已访问过的点标记为2
		countIsland(x+1, y)
		countIsland(x, y+1)
		countIsland(x-1, y)
		countIsland(x, y-1)
	}

	// 将边界0能达到的所有0都标记为1, 称为陆地沉没
	for i := 0; i < m; i++ {
		floodFill(i, 0)
		floodFill(i, n-1)
	}
	for j := 1; j < n-1; j++ {
		floodFill(0, j)
		floodFill(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 { // 该点是陆地，且没有被标记过
				// dfs遍历会将所有可达的点标记为2
				// 调用了dfs多少次，说明grid含有多少岛屿
				countIsland(i, j)
				res++
			}
		}
	}
	return res
}
// @lc code=end

