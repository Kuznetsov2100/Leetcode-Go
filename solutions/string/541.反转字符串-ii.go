/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	n := len(s)
	b := []byte(s)

	reverse := func(i, j int) {
		for ; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}

	for i := 0; i < n; i += 2*k {
		if i + k <= n { // 剩余字符 >= k个， 反转前k个字符
			reverse(i, i+k-1)
		} else {
			reverse(i, n-1) // 剩余字符 < k个， 反转剩下所有字符
		}
	}
	return string(b)
}
// @lc code=end

