/*
 * @lc app=leetcode.cn id=1260 lang=golang
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
// 本质是一维数组迁移，将一维信息映射到二维即可
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	size := m * n
	k = k % size
	if k == 0 {
		return grid
	}
	arr := make([]int, 0, size)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr = append(arr, grid[i][j])
		}
	}
	for i := 0; i < size; i++ {
		grid[i/n][i%n] = arr[(size - k + i) % size]
	}
	return grid
}



// 模拟法
// func shiftGrid(grid [][]int, k int) [][]int {
// 	m, n := len(grid), len(grid[0])
// 	k = k % (m*n)
// 	if k == 0 {
// 		return grid
// 	}
// 	for k > 0 {
// 		cache := grid[m-1][n-1]
// 		for i := m-1; i >= 0; i-- {
// 			for j := n-1; j > 0; j-- {
// 				grid[i][j] = grid[i][j-1]
// 			}
// 			if i > 0 {
// 				grid[i][0] = grid[i-1][n-1]
// 			}
// 		}
// 		grid[0][0] = cache
// 		k--
// 	}
// 	return grid
// }
// @lc code=end

