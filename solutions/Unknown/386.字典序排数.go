/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = i+1
	}
	sort.Slice(res, func(i, j int) bool {
		return strconv.Itoa(res[i]) < strconv.Itoa(res[j])
	})
	return res
}
// @lc code=end

