/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		num = num & (num-1) // n & (n-1)消除数字n的二进制表示中的最后一个1
		res++
	}
	return res
}
// @lc code=end

