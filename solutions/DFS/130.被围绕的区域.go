/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
// 本题给定的矩阵中有三种元素：
// 字母 X；
// 被字母 X 包围的字母 O；
// 没有被字母 X 包围的字母 O。
// 本题要求将所有被字母 X 包围的字母 O都变为字母 X ，但很难判断哪些 O 是被包围的，哪些 O 不是被包围的。
// 注意到题目解释中提到：任何边界上的 O 都不会被填充为 X。 我们可以想到，所有的不被包围的 O 都直接或间接与边界上的 O 相连。
// 我们可以利用这个性质判断 O 是否在边界上，具体地说：
// 对于每一个边界上的 O，我们以它为起点，标记所有与它直接或间接相连的字母 O；
// 最后我们遍历这个矩阵，对于每一个字母：
// 如果该字母被标记过，则该字母为没有被字母 X 包围的字母 O，我们将其还原为字母 O；
// 如果该字母没有被标记过，则该字母为被字母 X 包围的字母 O，我们将其修改为字母 X。
func solve(board [][]byte)  {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if (x < 0 || x >= m || y < 0 || y >= n) || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A' 
		dfs(x+1, y)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x, y-1)
	}

	// 将边界‘O’能达到的所有‘O’都标记为'A'
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
// @lc code=end

