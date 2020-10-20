 /*
 * @lc app=leetcode.cn id=807 lang=golang
 *
 * [807] 保持城市天际线
 */

// @lc code=start
func maxIncreaseKeepingSkyline(grid [][]int) int {
	var res int
	n := len(grid)
	rows := make([]int, n) // 保存每一行的最大值的数组
	cols := make([]int, n) // 保存每一列的最大值的数组
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rows[i] = max(rows[i], grid[i][j])
			cols[i] = max(cols[i], grid[j][i])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += min(rows[i], cols[j]) - grid[i][j]
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

