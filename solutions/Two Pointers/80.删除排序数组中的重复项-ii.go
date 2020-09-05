/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// 定义nums[0,index)区间内的每个元素最多出现两次 (核心思想：循环不变量)
	index := 0
	count := 1 // 记录当前数字出现的次数
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] != nums[i-1] {
				count = 1
			} else {
				count++
			}
		}
		if count <= 2 {
			nums[index] = nums[i]
			index++
		}
	}
	return index // index-1-0+1
}
// @lc code=end

