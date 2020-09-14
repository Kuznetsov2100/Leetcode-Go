/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */

// @lc code=start
func equationsPossible(equations []string) bool {
	// union-find算法
	// 遍历equations, 让 “a == b” 形式的方程左右两边的字母变量形成连通分量
	// 遍历equations，判断 "a != b" 形式的方程左右两边的字母变量是否连通，如果连通，那么逻辑矛盾
	uf := NewUF(26) // 26个小写英文字母
	for _, eq := range equations {
		if eq[1] == '=' {
			p := int(eq[0]-'a')
			q := int(eq[3]-'a')
			uf.Union(p, q)
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' {
			p := int(eq[0]-'a')
			q := int(eq[3]-'a')
			if uf.IsConnected(p, q) {
				return false
			}
		}
	}
	return true
}

type UF struct {
	parent []int
	size []int
	count int
}

func NewUF(n int) *UF {
	uf := &UF{}
	uf.count = n
	uf.parent = make([]int, n)
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (this *UF) IsConnected(p, q int) bool {
	return this.find(p) == this.find(q)
}

func (this *UF) Union(p, q int) {
	rootP, rootQ := this.find(p), this.find(q)
	if rootP == rootQ {
		return
	}
	if this.size[rootP] > this.size[rootQ] {
		this.parent[rootQ] = rootP
		this.size[rootP] += this.size[rootQ]
	} else {
		this.parent[rootP] = rootQ
		this.size[rootQ] += rootP
	}
	this.count--
}

func (this *UF) find(x int) int {
	for x != this.parent[x] {
		this.parent[x] = this.parent[this.parent[x]]
		x = this.parent[x]
	}
	return x
}
// @lc code=end

