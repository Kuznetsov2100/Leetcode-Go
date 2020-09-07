/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
import "bytes"
func findWords(board [][]byte, words []string) []string {
	var res []string
	m := len(board)
	if m == 0 || len(words) == 0 {
		return res
	}
	n := len(board[0])
	set := make(map[string]bool)
	prefix := make(map[string]bool)

	wordtrie := Constructor()
	for i := range words {
		wordtrie.Insert(words[i])
	}

	visited := make([][]bool, m) // 防止重复搜索
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	
	var dfs func(x, y int, word []byte) 
	dfs = func(x, y int, word []byte) {
		// 首先确保是棋盘内的合法坐标，其次该点没有访问过， 任一条件不满足，return
		if (x < 0 || x >= m || y < 0 || y >= n) || visited[x][y]  {
			return
		}
		word = append(word, board[x][y])
		str := string(word)
		if !prefix[str] {
			if !wordtrie.StartsWith(str) {
				return
			}
			prefix[str] = true
		}

		if wordtrie.Search(str) {
			if !set[str] {
				res = append(res, str)
				set[str] = true
			}
		}	

		visited[x][y] = true
		dfs(x, y+1, word)
		dfs(x+1, y, word)
		dfs(x, y-1, word)
		dfs(x-1, y, word)	
		visited[x][y] = false	
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, []byte{})
		}
	}
	return res
}

type Trie struct {
	root *node
}

type node struct {
	isString bool
	next [26]*node
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	this.root = this.insert(this.root, word, 0)
}

func (this *Trie) insert(x *node, key string, d int) *node {
	if x == nil {
		x = &node{}
	}
	if d == len(key) {
		x.isString = true
	} else {
		x.next[key[d]-97] = this.insert(x.next[key[d]-97], key, d+1)
	}
	return x
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	x := this.get(this.root, word, 0)
	if x == nil {
		return false
	}
	return x.isString
}

func (this *Trie) get(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	return this.get(x.next[key[d]-97], key, d+1)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var b bytes.Buffer
	b.WriteString(prefix)
	return this.collectPrefix(this.get(this.root, prefix, 0), &b)
}

func (this *Trie) collectPrefix(x *node, prefix *bytes.Buffer) bool {
	if x == nil {
		return false
	}	
	if x.isString {
		return true
	}
	for c := 97; c <= 122; c++ {
		prefix.WriteByte(byte(c))
		if this.collectPrefix(x.next[c-97], prefix) {
			return true
		}
		prefix.Truncate(prefix.Len()-1)
	}
	return false
}
// @lc code=end

