/*
 * @lc app=leetcode.cn id=685 lang=golang
 *
 * [685] 冗余连接 II
 */

// @lc code=start
func findRedundantDirectedConnection(edges [][]int) []int {
	// 首先遍历edges,判断是否有某个节点的入度为2，如果没有，则按684题的方法解答
	// 如果存在节点A的入度为2(有且只有一个入度为2的节点),例如，B->A,C->A, 记录节点B，C
	// 倒序遍历edges,跳过边(C->A),其他所有边union,如果Count为2(节点0算一个连通分量),
	// 说明删除边(C->A)后，剩下的图是有N个节点的有根树则边(C->A)是正确答案
	// 如果Count>2, 删除(C->A)后， 剩下的图无法形成N个节点的有根树，则边(B->A)是正确答案
	var indegree2, first, second int
	N := len(edges)
	indegree := make(map[int][]int)
	for _, edge := range edges {
		indegree[edge[1]] = append(indegree[edge[1]], edge[0])
		if len(indegree[edge[1]]) == 2 {
			indegree2 = edge[1]
			first, second = indegree[edge[1]][0], indegree[edge[1]][1]
			break
		}
	}
	
	if indegree2 == 0 {
		var res []int
		uf := NewUF(N+1)
		for _, edge := range edges {
			u, v := edge[0], edge[1]
			if uf.IsConnected(u,v) {
				res = []int{u, v}
			}
			uf.Union(u, v)
		}
		return res
	}

	uf1 := NewUF(N+1)
	for i := N-1; i >= 0; i-- {
		u, v := edges[i][0], edges[i][1]
		if v == indegree2 && u == second {
			continue
		}
		uf1.Union(u,v)
	}

	if uf1.Count() == 2 {
		return []int{second, indegree2}
	}
	return []int{first, indegree2}
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

