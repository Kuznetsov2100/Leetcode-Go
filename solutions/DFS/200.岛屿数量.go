/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 坐标（x,y）非法 或者是水域或者是已访问过的点，则返回
		if (x < 0 || x >= m || y < 0 || y >= n) || grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2' // 已访问过的点标记为2
		dfs(x+1, y)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x, y-1)
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' { // 该点是陆地，且没有被标记过
				// dfs遍历会将所有可达的点标记为2
				// 调用了dfs多少次，说明grid含有多少岛屿
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
// @lc code=end

