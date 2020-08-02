/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
import "sort"
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	used := make([]bool, len(nums))
	var bt func(path []int)
	bt = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp,path)		
			res = append(res, tmp)
			return
		}
		for i := range nums {
			if !used[i] { // 访问未访问的
				if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
					continue
				}
				path = append(path, nums[i]) // 做选择
				used[i] = true
				bt(path) // 进入下一层决策树
				path = path[:len(path)-1] // 撤销选择
				used[i] = false
			}
		}
	}
	bt([]int{})
	return res
}
// @lc code=end

