/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	isValid := func(r, c int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[r][i] == num && i != c {
				return false
			}
			if board[i][c] == num && i != r {
				return false
			}
			if row, col := ((r/3)*3+i/3),((c/3)*3+i%3); board[row][col] == num && row != r && col != c  {
				return false
			}
		}
		return  true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !isValid(i,j,board[i][j]) {
					return false
				}
			}
		}
	}
	return true	
}

// 由嵌套循环的知识(外层固定，内层遍历)：
// borad[i][j]表示了对固定的一行i，遍历列j(即每次检查一整行)
// board[j][i]表示了对固定的一列i，遍历行j(即每次检查一整列)
// board[(i/3)*3+j/3][(i%3)*3+j%3]表示了对固定的宫i，遍历格j(即每次检查一个宫)
// func isValidSudoku(board [][]byte) bool {
// 	for i := 0; i < 9; i++ {
// 		mapr := make(map[byte]bool)
// 		mapc := make(map[byte]bool)
// 		maps := make(map[byte]bool)
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] != '.' {
// 				if mapr[board[i][j]] {
// 					return false
// 				}
// 				mapr[board[i][j]] = true
// 			}
// 			if board[j][i] != '.' {
// 				if mapc[board[j][i]] {
// 					return false
// 				}
// 				mapc[board[j][i]] = true
// 			}
// 			if row, col := ((i/3)*3+j/3),((i%3)*3+j%3); board[row][col] != '.' {
// 				if maps[board[row][col]] {
// 					return false
// 				}						
// 				maps[board[row][col]] = true
// 			}
// 		}
// 	}
// 	return true
// }
// @lc code=end

