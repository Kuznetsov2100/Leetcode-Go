/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	// dfs + 记忆化
	var res int
	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右偏移量数组
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if memo[x][y] != 0 { // memo[x][y] != 0 说明以坐标(x,y)为起点的最长递增路径长度已计算出来
			return memo[x][y]
		}
		memo[x][y] = 1 // 初始化长度为1
		for i := 0; i < 4; i++ {
			newx := x+directions[i][0]
			newy := y+directions[i][1]
			if newx >= 0 && newx < m && newy >= 0 && newy < n && matrix[newx][newy] > matrix[x][y] {
				memo[x][y] = max(memo[x][y], dfs(newx, newy)+1)
			}
		}
		return memo[x][y] // (x,y)为起点的最长递增路径长度
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

