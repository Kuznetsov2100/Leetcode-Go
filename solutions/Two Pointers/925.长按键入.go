/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	// 双指针
	i, j, m, n := 0, 0, len(name), len(typed)
	for i < m && j < n {
		char1, char2 := name[i], typed[j]
		if char1 != char2 {
			return false
		}
		cnt1, cnt2 := 0, 0
		for i < m && name[i] == char1 {
			i++
			cnt1++
		}
		for j < n && typed[j] == char2 {
			j++
			cnt2++
		}
		if cnt2 < cnt1 {
			return false
		}
	}
	return i == m && j == n
}
// @lc code=end

