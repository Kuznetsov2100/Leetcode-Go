/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var res [][]int
	if n <= 0 || k <= 0 {
		return res
	}
	var backtrack func(start int, track []int)
	backtrack = func(start int, track []int) {
		if len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i+1, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(1, []int{})
	return res
}
// @lc code=end

