/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	var res [][]string
	// 初始化棋盘
	queenCol := make([]int,n)
	for i := range queenCol {
		queenCol[i] = -1
	}
	
	// 路径：board中小于row的那些行都已经成功放置了皇后
	// 可选择列表: 第row行的所有列都是放置Q的选择
    // 结束条件: row超过board的最后一行
	var backtrack func(row int, queenCol []int)
	backtrack = func(row int,  queenCol []int) {
		if row == n { // 皇后放置完毕
			var tmp []string
			board := make([][]byte,n)
			for i := range board {
				board[i] = make([]byte, n)
				for j := range board[i] {
					board[i][j] = '.'
				}
				board[i][queenCol[i]] = 'Q'
				tmp = append(tmp,string(board[i]))
			}
			res = append(res, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(queenCol, row, col) {
				continue
			}
			queenCol[row] = col
			backtrack(row+1, queenCol)
			queenCol[row] = -1
		}
	}
	backtrack(0, queenCol)
	return res
}
func isValid(queenCol []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if col == queenCol[i] || abs(row - i) == abs(col - queenCol[i]) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

