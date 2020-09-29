/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	var res int
	n, count := len(arr), 1
	threshold := n/4
	for i := 0; i < n-1; i++ {
		if arr[i] != arr[i+1] {
			if count > threshold {
				res = arr[i]
				break
			} else {
				count = 1
				continue
			}
		}
		count++
	}
	if res == 0 && count > threshold {
		res = arr[n-1]
	}
	return res
}
// @lc code=end

