/*
 * @lc app=leetcode.cn id=1460 lang=golang
 *
 * [1460] 通过翻转子数组使两个数组相等
 */

// @lc code=start
func canBeEqual(target []int, arr []int) bool {
	bucket := make([]int, 1001)
	for i := range target {
		bucket[target[i]]++
		bucket[arr[i]]--
	}
	for _, num := range bucket {
		if num != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

