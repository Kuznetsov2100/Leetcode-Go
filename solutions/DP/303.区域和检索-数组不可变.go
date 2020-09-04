/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	dp []int
}


func Constructor(nums []int) NumArray {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[0]
			continue
		}
		dp[i] = dp[i-1]+nums[i]
	}
	return NumArray{dp: dp}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.dp[j]
	}
	return this.dp[j]-this.dp[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

