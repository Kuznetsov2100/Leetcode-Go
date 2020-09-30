/*
 * @lc app=leetcode.cn id=1525 lang=golang
 *
 * [1525] 字符串的好分割数目
 */

// @lc code=start
func numSplits(s string) int {
	n := len(s)
	dpLeft := make([]int, n+1)
	dpRight := make([]int, n+1)
	setLeft := make([]bool, 26)
	setRight := make([]bool, 26)
	for i := 1; i <= n; i++ {
		if setLeft[s[i-1]-'a'] {
			dpLeft[i] = dpLeft[i-1]
		} else {
			dpLeft[i] = dpLeft[i-1] + 1
			setLeft[s[i-1]-'a'] = true
		}
	}
	
	for i := n-1; i >= 0; i-- {
		if setRight[s[i]-'a'] {
			dpRight[n-i] = dpRight[n-i-1] 
		} else {
			dpRight[n-i] = dpRight[n-i-1] + 1
			setRight[s[i]-'a'] = true
		}
	}

	
	var res int
	for i := 1; i < n; i++ {
		if dpLeft[i] == dpRight[n-i] {
			res++
		}
	}
	return res
}
// @lc code=end

