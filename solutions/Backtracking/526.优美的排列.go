/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 */

// @lc code=start
func countArrangement(N int) int {
	var res int
	used := make([]bool, N)
	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == N+1 {
			res++
			return
		}
		for i := 1; i <= N; i++ {
			if !used[i-1] && (pos % i == 0 || i % pos == 0) {		
				used[i-1] = true
				backtrack(pos+1) // 进入下一层决策树
				used[i-1] = false			
			}
		}
	}
	backtrack(1)
	return res
}
// @lc code=end

