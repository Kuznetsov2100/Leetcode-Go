/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
import "sort"
func subsetsWithDup(nums []int) [][]int {
	// 需要对nums排序
	sort.Ints(nums)
	var res [][]int
	used := make([]bool, len(nums))
	var backtrack func(start int, track []int)
	backtrack = func(start int, track []int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			if !used[i] {
				// 剪枝
				if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
					continue
				}
				track = append(track, nums[i])
				used[i] = true
				backtrack(i+1, track)
				track = track[:len(track)-1]
				used[i] = false
			}
		}
	}
	backtrack(0, []int{})
	return res
}
// @lc code=end

