/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var backtrack func(start int, path []int, pre int, length int)
	backtrack = func(start int, path []int, pre int, length int) {
		if length > 1 {
			tmp := make([]int, length)
			copy(tmp,path)
			res = append(res, tmp)
		}
		cache := make(map[int]int)
		for i := start; i < len(nums); i++ {
			if _, ok := cache[nums[i]]; nums[i] < pre || ok {
				continue
			}
			path = append(path, nums[i])
			cache[nums[i]] = 1
			backtrack(i+1, path, nums[i], length+1)
			path = path[:len(path)-1]		
		}
	}
	backtrack(0, []int{}, -101, 0)
	return res
}
// @lc code=end

