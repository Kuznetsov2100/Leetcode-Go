/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	var res []int
	uf := NewUF(len(edges)+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if uf.IsConnected(u,v) {
			res = []int{u, v}
		}
		uf.Union(u, v)
	}
	return res
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

func (this *UF) Count() int {
	return this.count
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

