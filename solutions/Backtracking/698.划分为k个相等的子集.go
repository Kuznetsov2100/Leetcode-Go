/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	var totalSum int
	for i := range nums {
		totalSum += nums[i]
	}
	subSum := totalSum/k
	if totalSum % k != 0 { // k无法整除totalSum
		return false
	}
	// 数组按降序排序，凑数的时候，从左往右，先考虑大数，达到剪枝的目的
	sort.Slice(nums, func(i, j int) bool {return nums[i] > nums[j]})
	if nums[0] > subSum { // 数组最大数大于子集的和
		return false
	}
	numsLength := len(nums)
	used := make([]bool, numsLength) // 避免重复搜索
	var backtracking func(start, sum, k int) bool
	backtracking = func(start, sum, k int) bool {
		if k == 1 { // 已找到了k-1个子集，剩下的数的和一定等于subSum
			return true
		}
		if sum > subSum {
			return false
		}
		if sum == subSum {
			return backtracking(0, 0, k-1) // 凑满了一个子集，开始凑下一个子集
		}
		for i := start; i < numsLength; i++ {
			if !used[i] {
				used[i] = true
				if backtracking(i+1, sum+nums[i], k) {
					return true // 可以将数组划分为k个和相等的非空子集
				}
				used[i] = false // 回撤
			}
		}
		return false
	}
	return backtracking(0, 0, k)
}
// @lc code=end

