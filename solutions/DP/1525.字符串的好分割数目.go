/*
 * @lc app=leetcode.cn id=1525 lang=golang
 *
 * [1525] 字符串的好分割数目
 */

// @lc code=start
func numSplits(s string) int {
	n := len(s)
	dp1 := make([]int, n+1)
	set1 := make(map[byte]bool)
	for i := 1; i <= n; i++ {
		if set1[s[i-1]] {
			dp1[i] = dp1[i-1]
		} else {
			dp1[i] = dp1[i-1] + 1
			set1[s[i-1]] = true
		}
	}
	

	dp2 := make([]int, n+1)
	set2 := make(map[byte]bool)
	for i := n-1; i >= 0; i-- {
		if set2[s[i]] {
			dp2[n-i] = dp2[n-i-1] 
		} else {
			dp2[n-i] = dp2[n-i-1] + 1
			set2[s[i]] = true
		}
	}

	
	var res int
	for i := 0; i < n; i++ {
		if dp1[i+1] == dp2[n-i-1] {
			res++
		}
	}
	return res
}
// @lc code=end

