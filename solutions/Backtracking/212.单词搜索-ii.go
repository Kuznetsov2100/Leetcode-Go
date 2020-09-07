/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
// 思路： 使用字典树(trie)的回溯
// step1 : 根据字典中的单词构建一个trie
// step2 : 对board中每个单元格使用回溯搜索
// 在dfs中，检查到目前为止遍历的字母序列是否是字典树中任一单词的前缀，
// 可以及时停止回溯
func findWords(board [][]byte, words []string) []string {
	var res []string
	m := len(board)
	if m == 0 || len(words) == 0 {
		return res
	}
	n := len(board[0])
	wordtrie := Constructor()
	for i := range words {
		wordtrie.Insert(words[i])
	}

	visited := make([][]bool, m)  // 防止重复搜索
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	
	var dfs func(x, y int, node *Trie) 
	dfs = func(x, y int, node *Trie) {
		// 首先确保是棋盘内的合法坐标，其次该点没有访问过， 任一条件不满足,  return
		if (x < 0 || x >= m || y < 0 || y >= n) || visited[x][y]  {
			return
		}
		node = node.children[board[x][y]-'a']
		if node == nil {
			return
		}
		if node.isEnd {
			res = append(res, node.str)
			node.isEnd = false // 防止重复添加单词
		}
		visited[x][y] = true
		dfs(x, y+1, node)
		dfs(x+1, y, node)
		dfs(x, y-1, node)
		dfs(x-1, y, node)	
		visited[x][y] = false	
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, &wordtrie)
		}
	}
	return res
}


type Trie struct {
	isEnd bool
	str string
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for _, ch := range word {
		index := ch-'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{} 
		}
		node = node.children[index]
	}
	node.str =  word
	node.isEnd = true
}
// @lc code=end

