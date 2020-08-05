/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
import "sort"
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	var backtrack func(sum int, track []int)
	backtrack = func(sum int, track []int) {
		if sum == target && sort.IntsAreSorted(track) {
			tmp := make([]int, len(track))
			copy(tmp, track)	
			res = append(res, tmp)
			return
			
			
		}
		for i := 0; i < len(candidates); i++ {		
			sum += candidates[i]
			if sum > target {
				continue
			}
			track = append(track, candidates[i])
			backtrack(sum, track)
			track = track[:len(track)-1]
			sum -= candidates[i]					
		}
	}
	backtrack(0,[]int{})
	return res
}
// @lc code=end

