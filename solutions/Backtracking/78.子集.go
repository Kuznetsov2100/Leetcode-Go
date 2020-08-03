/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var res [][]int
	var backtrack func(start int, track []int)
	backtrack = func(start int, track []int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i+1, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(0, []int{})
	return res
}
// @lc code=end

