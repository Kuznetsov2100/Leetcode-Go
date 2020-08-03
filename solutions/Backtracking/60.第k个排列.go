/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	var res string
	var count int
	used := make([]bool, n)
	var bt func(path string)
	bt = func(path string) {
		if len(path) == n {
			count++
			if count == k {
				res = path
			}
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i-1] { // 访问未访问的
				if res != "" {
					return
				}
				path += strconv.Itoa(i) // 做选择
				used[i-1] = true
				bt(path) // 进入下一层决策树
				path = path[:len(path)-1] // 撤销选择
				used[i-1] = false
			}
		}
	}
	bt("")
	return res
}
// @lc code=end

