/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	isValid := func(r, c int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[r][i] == num { // 该行是否已填写num
				return false
			}
			if board[i][c] == num { // 该列是否已填写num
				return false
			}
			if board[(r/3)*3+i/3][(c/3)*3+i%3] == num { // 该格子所处的3X3方框是否已填写num
				return false
			}
		}
		return  true
	}
	var backtrack func(r, c int) bool
	backtrack = func(r, c int) bool {
		if c == 9 { // 穷举到当行的最后一列后，切换到下一行的第一列开始
			return backtrack(r+1,0)	
		}
		if r == 9 { // 找到一个可行解
			return true
		}
		if board[r][c] != '.' { // 此格子已有数字，直接跳到下一列
			return backtrack(r, c+1)
		}
		for num := byte('1'); num <= '9'; num++ {	
			if !isValid(r, c, num) { // 跳过不合法的数字
				continue
			}
			board[r][c] = num
			if backtrack(r, c+1) {
				return true // 找到一个可行解，直接结束
			}
			board[r][c] = '.'
		}
		return false // 穷举结束后，依然没有得到可行解
	}
	backtrack(0,0)
}
// @lc code=end

