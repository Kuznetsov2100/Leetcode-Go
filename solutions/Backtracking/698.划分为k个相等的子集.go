/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	var totalSum, maxNum int
	for i := range nums {
		totalSum += nums[i]
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	subSum := totalSum/k
	if totalSum % k != 0 || maxNum > subSum {
		return false
	}
	numsLength := len(nums)
	used := make([]bool, numsLength)
	var backtracking func(start, sum, k int) bool
	backtracking = func(start, sum, k int) bool {
		if k == 0 {
			return true
		}
		if sum > subSum {
			return false
		}
		if sum == subSum {
			return backtracking(0, 0, k-1)
		}
		for i := start; i < numsLength; i++ {
			if !used[i] {
				used[i] = true
				if backtracking(i+1, sum+nums[i], k) {
					return true
				}
				used[i] = false
			}
		}
		return false
	}
	return backtracking(0, 0, k)
}
// @lc code=end

