/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */

// @lc code=start
func findCircleNum(M [][]int) int {
	// union-find算法
	// 如果有M[i][j]=1,那么M[j][i]=1
	// 因此只需要遍历矩阵的下三角即可(不包含对角线上的点)
	N := len(M)
	uf := NewUF(N) // N个学生
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j) // 学生i和学生j建立朋友关系
			}
		}
	}
	return uf.Count // 朋友圈总数
}

type UF struct {
	parent []int
	size []int
	Count int
}

func NewUF(n int) *UF {
	uf := &UF{}
	uf.Count = n
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
	this.Count--
}

func (this *UF) find(x int) int {
	for x != this.parent[x] {
		this.parent[x] = this.parent[this.parent[x]]
		x = this.parent[x]
	}
	return x
}
// @lc code=end

