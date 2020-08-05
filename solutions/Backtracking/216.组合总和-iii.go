/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	if k <= 0 || n <= 0 {
		return res
	}
	var backtrack func(start,sum, pre int, track []int)
	backtrack = func(start,sum, pre int, track []int) {
		if sum == n && len(track) == k { 
			tmp := make([]int, len(track))
			copy(tmp, track)	
			res = append(res, tmp)
			return		
		}
		for i := start; i <= 9; i++ {
			if i <= pre {
				continue
			}
			sum += i
			if sum > n {
				continue
			}
			track = append(track, i)
			backtrack(start+1,sum,i,track)
			track = track[:len(track)-1]
			sum -= i					
		}
	}
	backtrack(1,0,-101, []int{})
	return res
}
// @lc code=end

