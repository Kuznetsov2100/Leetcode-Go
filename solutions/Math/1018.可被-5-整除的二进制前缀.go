/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	mod5remainder := []int{0,1,2,3,4,0,1,2,3,4}
	num := 0
	for i, digit := range A {
		num = mod5remainder[num<<1+digit]
		if  num == 0 {
			res[i] = true
		}
	}
	return res
}
// @lc code=end

