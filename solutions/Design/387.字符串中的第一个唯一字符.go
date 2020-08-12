/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
import "strings"
func firstUniqChar(s string) int {
	arr := make([]int, 26)
	for i := range s {
		arr[s[i]-97]++
	}
	index := 1 << 63 -1
	for i := range arr {
		if arr[i] == 1 {
			index = min(index, strings.Index(s,string(byte(i+97))))
		}
	}
	if index == (1 << 63 -1) {
		return -1
	}
	return index
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

