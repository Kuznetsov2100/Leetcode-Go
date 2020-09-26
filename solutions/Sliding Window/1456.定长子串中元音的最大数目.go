/*
 * @lc app=leetcode.cn id=1456 lang=golang
 *
 * [1456] 定长子串中元音的最大数目
 */

// @lc code=start
func maxVowels(s string, k int) int {
	// 滑动窗口， 窗口长度为k, count为当前窗口中含有的元音字母数量
	var res int
	left, right, count := 0, 0, 0
	for right < len(s) {
		char := s[right]
		if isVowel(char) {
			count++
		}
		right++
		if right - left + 1 > k { // 窗口长度大于k,需要缩小窗口
			res = max(res, count) // 更新res
			char := s[left]
			if isVowel(char) {
				count--
			}
			left++
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

func isVowel(x byte) bool {
	if x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' {
		return true
	}
	return false
}
// @lc code=end

