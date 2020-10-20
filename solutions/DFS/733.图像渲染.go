/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// 洪水填充算法， 其思想是利用dfs可达性来更新所有可达的位置的颜色
	// 如果initialColor 与 newColor相同，则直接返回image
	// 本题无需visited数组来避免重复访问，因为newColor 与 initialColor 不同
	initialColor := image[sr][sc]
	if initialColor == newColor {
		return image
	}
	m, n := len(image), len(image[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if image[x][y] == initialColor {
			image[x][y] = newColor
			dfs(x+1, y)
			dfs(x, y+1)
			dfs(x-1, y)
			dfs(x, y-1)
		}
	}
	dfs(sr, sc)
	return image
}
// @lc code=end

