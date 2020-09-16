/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
// 本题与基本的union-find模板有很大的不同，需要自定义数据结构以及正确处理权值关系。
// 假设 a/b = 2, 可以令parent[a] = b，即a的父节点为b, ratio[a] = 2,
// 表明a与它的父节点b的比值为2. 在遍历equations的时候，要考虑以下情况：
// case1: a, b 都不存在于任何集合中，此时令parent[a] = b, ratio[a] = value[i], parent[b] = b, ratio[b] = 1.0
// case2:b存在于某个集合中，a不存在于任何集合中，令parent[a] = b, ratio[a] = value[i]
// case3:a存在与某个集合中，b不存在于任何集合中，令parent[b] = a, ratio[b] = 1.0 / values[i]
// case4:a,b都已加入集合，此时对a,b进行union操作，下面解释具体的union操作：
// 情况1： a,b都属于同一集合，那么不需要进行任何操作
// 情况2: a,b属于不同的集合，我们需要对这两个集合合并
// 在合并集合之前，先解释一下find函数的作用：
// find函数除了返回变量x的根节点，还进行了路径压缩以及权值的更新, 非根节点直接指向根节点，变量与父节点的比值变成了变量与根节点的比值。

// 下面讲解如何合并集合：
// 假设a的根节点为root1,b的根节点为root2，让root2成为root1的父节点，即parent[root1] = root2,
// ratio[root1]的值本来是1.0,即根节点与它的父节点(它自己) root1/root1 = 1.0,现在root1的父节点
// 变成了root2,我们需要更新ratio[root1]的值, 即root1/root2 = ?,可以利用已知信息来求：
// 我们有 ratio[b] = b/root2, ratio[a] = a/root1，以及a/b = value
// root1/root2 = b/root2 * a/b / a/root1
// 因此 ratio[root1] = root1/root2 = value*ratio[b]/ratio[a]
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	uf := NewUF()
	for i, value := range values {
		from, to := equations[i][0], equations[i][1]
		_, okfrom := uf.Parent[from]
		_, okto := uf.Parent[to]
		if !okfrom && !okto {
			uf.Parent[from] = to
			uf.Ratio[from] = value
			uf.Parent[to] = to
			uf.Ratio[to] = 1.0
			continue
		}
		if !okfrom {
			uf.Parent[from] = to
			uf.Ratio[from] = value
			continue
		}
		if !okto {
			uf.Parent[to] = from
			uf.Ratio[to] = 1.0 / value
			continue
		}
		uf.Union(from, to, value)
	}

	var res []float64
	for _, query := range queries {
		from, to := query[0], query[1]
		_, okfrom := uf.Parent[from]
		_, okto := uf.Parent[to]
		if !okfrom || !okto { // from, to 有一个不属于任何集合，则无法求值
			res = append(res, -1.0)
			continue
		}
		if uf.Find(from) != uf.Find(to) { //from和to不属于同一集合，无法求值
			res = append(res, -1.0)
			continue
		}
		res = append(res, uf.Ratio[from] / uf.Ratio[to])
	}
	return res
}

type UF struct {
	Parent map[string]string
	Ratio map[string]float64
}

func NewUF() *UF {
	return &UF{
		Parent: make(map[string]string),
		Ratio: make(map[string]float64),
	}
}

// 查找变量x所属集合的根节点
func (this *UF) Find(x string) string {
	if x != this.Parent[x] {
		px := this.Parent[x]
		this.Parent[x] = this.Find(px) // x指向根节点
		this.Ratio[x] *= this.Ratio[px] // 更新x与根节点的比值 ratio[x] = x/px, ratio[px] = px/root, x/root = x/px * px/root
	}
	return this.Parent[x]
}

func (this *UF) Union(from, to string, ratio float64) {
	pfrom := this.Find(from)
	pto := this.Find(to)
	if pfrom == pto {
		return
	}
	this.Parent[pfrom] = pto
	this.Ratio[pfrom] = ratio * this.Ratio[to] / this.Ratio[from]
}
// @lc code=end

