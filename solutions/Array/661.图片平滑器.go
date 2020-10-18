/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */

// @lc code=start
func imageSmoother(M [][]int) [][]int {
	row, col := len(M), len(M[0])
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, col)
	}
	inArea := func(r, c int) bool {
		if r < 0 || c < 0 || r >= row  || c >= col {
			return false
		}
		return true
	}
	countGrayScale := func(r, c int) int {
		sum, count := 0, 0
		for i := 0; i < 9; i++ {
			x, y := r-1+i/3, c-1+i%3
			if inArea(x, y) {
				count++
				sum += M[x][y]
			}
		}
		return sum/count
	}
	
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[i][j] = countGrayScale(i, j)
		}
	}
	return res
}
// @lc code=end

