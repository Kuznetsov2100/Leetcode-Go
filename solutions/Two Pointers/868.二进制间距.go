/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(N int) int {
	var res int
	binN := strconv.FormatInt(int64(N), 2)
	slow, fast := 0, 1
	for fast < len(binN) {
		if binN[fast] == '1' {
			res = max(res, fast-slow)
			slow = fast
		}
		fast++
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

