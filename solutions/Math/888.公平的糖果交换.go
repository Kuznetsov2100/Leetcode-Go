/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
func fairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	var res []int
	m := make(map[int]bool)
	for _, num := range A {
		sumA += num
	}
	for _, num := range B {
		sumB += num
		m[num] = true
	}
	delta := (sumB-sumA)/2
	for _, num := range A {
		if m[num+delta] {
			res = []int{num, num+delta}
			break
		}
	}
	return res
}
// @lc code=end

