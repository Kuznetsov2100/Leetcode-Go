/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
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

