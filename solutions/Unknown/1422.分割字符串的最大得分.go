/*
 * @lc app=leetcode.cn id=1422 lang=golang
 *
 * [1422] 分割字符串的最大得分
 */

// @lc code=start
func maxScore(s string) int {
	var res int
	for i := 1; i < len(s); i++ {
		res = max(res, count0(s[:i])+count1(s[i:]))
	}
	return res
}

func count0(s string) int {
	var res int
	for i := range s {
		if s[i] == '0' {
			res++
		}
	}
	return res
}

func count1(s string) int {
	var res int
	for i := range s {
		if s[i] == '1' {
			res++
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

