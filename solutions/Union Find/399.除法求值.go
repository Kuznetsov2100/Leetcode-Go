/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
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
		if !okfrom || !okto {
			res = append(res, -1.0)
			continue
		}
		if uf.Find(from) != uf.Find(to) {
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

func (this *UF) Find(x string) string {
	if x != this.Parent[x] {
		px := this.Parent[x]
		this.Parent[x] = this.Find(px)
		this.Ratio[x] *= this.Ratio[px]
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

