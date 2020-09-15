/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	var res [][]string
	email_id := make(map[string]int) // 把每个邮箱映射到所在的账户在accounts数组中的索引(账户id)
	accountid_emails := make(map[int][]string) // 每个账户id所对应的邮箱集合
	uf := NewUF(len(accounts)) // 初始化union-find对象
	for i, account := range accounts {
		for _, email := range account[1:] {
			id, ok := email_id[email]
			if !ok {
				email_id[email] = i // 建立email -> id映射关系
			} else {
				uf.Union(id, i) // 合并账户
			}
		}
	}	
	for email, id := range email_id {
		accountid := uf.Find(id) // 查找该账户id所属的账户集合(连通分量)，两个账户id合并后，则属于同一个集合
		accountid_emails[accountid] = append(accountid_emails[accountid], email)
	}	
	for accountid, emails := range accountid_emails {
		sort.Slice(emails, func(i, j int) bool {return emails[i] < emails[j]}) // email按升序排序
		res = append(res, append([]string{accounts[accountid][0]}, emails...)) // 拼接结果
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
	return this.Find(p) == this.Find(q)
}

func (this *UF) Union(p, q int) {
	rootP, rootQ := this.Find(p), this.Find(q)
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

func (this *UF) Find(x int) int {
	for x != this.parent[x] {
		this.parent[x] = this.parent[this.parent[x]]
		x = this.parent[x]
	}
	return x
}
// @lc code=end

