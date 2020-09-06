/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	wordLength := len(word)
	visited := make([][]bool, m) // 防止重复搜索
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		// 首先确保是棋盘内的合法坐标，其次该点没有访问过，最后该点字母等于word相应索引处字母， 任一条件不满足，return false
		if (x < 0 || x >= m || y < 0 || y >= n) || visited[x][y] || board[x][y] != word[index] {
			return false
		}

		if index == wordLength-1 { // 单词存在于网格中
			return true
		}

		visited[x][y] = true
		if dfs(x, y+1, index+1) || dfs(x+1, y, index+1) || dfs(x, y-1, index+1) || dfs(x-1, y, index+1) {
			return true
		}	
		visited[x][y] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
// @lc code=end

