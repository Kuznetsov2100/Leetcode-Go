/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

// @lc code=start
func duplicateZeros(arr []int)  {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
}
// @lc code=end

