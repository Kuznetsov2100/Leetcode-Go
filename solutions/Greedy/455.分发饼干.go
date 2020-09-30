/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	// 先排序，然后贪心思想
	// 从小到大开始遍历两个数组：
	// 如果当前s[j] >= g[i]，则可以满足该孩子的胃口，res++，i++,j++。
	// 否则，j++，看下一个饼干的大小能否满足对应孩子的胃口。

	var res int
	m, n := len(g), len(s)
	if m == 0 || n == 0 {
		return res
	}
	sort.Ints(g)
	sort.Ints(s)
	
	i, j := 0, 0
	for i < m && j < n {
		if s[j] >= g[i] {
			res++
			i++
			j++
		} else {
			j++
		}
	}
	return res
}
// @lc code=end

