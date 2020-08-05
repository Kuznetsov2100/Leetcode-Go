/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
import "sort"
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	used := make([]bool, len(candidates))
	var backtrack func(start, sum int, track []int)
	backtrack = func(start, sum int, track []int) {
		if sum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if !used[i] {
				// 剪枝
				if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
					continue
				}
				sum += candidates[i]
				if sum > target {
					continue
				}
				track = append(track, candidates[i])
				used[i] = true
				backtrack(i+1, sum, track)
				track = track[:len(track)-1]
				sum -= candidates[i]
				used[i] = false
			}
		}
	}
	backtrack(0,0,[]int{})
	return res
}
// @lc code=end

