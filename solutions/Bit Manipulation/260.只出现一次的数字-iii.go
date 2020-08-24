/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	// 设两个出现次数为1次的数分别为x,y
	// 对数组所有数异或，bitmap为 x ^ y的结果
	// 获取bitmap最右边出现的1,这个1表示x和y在这一位不同
	// 对该位置上为1的数进行异或，结果为x，y 为 x ^ bitmap 的结果
	bitmap := 0
	for _, num := range nums {
		bitmap ^= num
	}
	diff := bitmap & (-bitmap) 
	x := 0
	for _, num := range nums {
		if num & diff != 0 {
			x ^= num
		}
	}
	return []int{x, x ^ bitmap}
}
// @lc code=end

