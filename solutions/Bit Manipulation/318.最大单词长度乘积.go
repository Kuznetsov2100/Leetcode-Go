/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	// 题目中单词只包含小写字母，将26个字母映射到26位二进制，1表示该位对应的字母出现在单词中，0表示没有出现。
	// 如果两个单词没有公共字母，则两个单词的二进制表示的 bitwise and 结果为0，否则为1。
	res := 0
	n := len(words)
	lens := make([]int, n)
	bitmask := make([]int, n)
	for i := 0; i < n; i++ {
		bit := 0
		for _, ch := range words[i] {
			bit |= 1 << (ch-'a') // 将字母所对应的二进制位设为1
		}
		bitmask[i] = bit // bit即为该单词的二进制表示
		lens[i] = len(words[i])
	}
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if bitmask[i] & bitmask[j] == 0 { // 两个单词没有公共字母
				if length := lens[i]*lens[j]; length > res {
					res = length
				}
			}
		}
	}
	return res
}
// @lc code=end

