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
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]+nums[i]
	}
	return NumArray{dp: dp}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j+1]-this.dp[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

