/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	// 假设数组长度为n,答案范围为[1,n+1]
	// 原地哈希，将nums[i]映射到nums[i]-1位置上,对号入座
	// 第二次遍历数组，检查每个数字是否按照要求入座，如果号座不一致，即找到目标
	n := len(nums)
	for i := 0; i < n; i++ {
		for 1 <= nums[i] && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i+1
		}
	}
	return n+1 // 数组1..n都出现了，那么未出现的最小正整数即是n+1
}
// @lc code=end

