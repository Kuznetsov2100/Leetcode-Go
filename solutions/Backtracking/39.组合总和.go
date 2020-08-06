/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
import "sort"
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var backtrack func(sum, pre int, track []int)
	backtrack = func(sum, pre int, track []int) {
		if sum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)	
			res = append(res, tmp)
			return		
		}
		for i := 0; i < len(candidates); i++ {
			if candidates[i] < pre {
				continue
			}
			sum += candidates[i]
			if sum > target {
				continue
			}
			track = append(track, candidates[i])
			backtrack(sum,candidates[i],track)
			track = track[:len(track)-1]
			sum -= candidates[i]					
		}
	}
	backtrack(0,-101, []int{})
	return res
}
// @lc code=end

