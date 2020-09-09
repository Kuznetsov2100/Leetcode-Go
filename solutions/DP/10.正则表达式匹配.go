/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
// 以一个例子详解动态规划转移方程：
// S = abbbbc
// P = ab*d*c
// 1. 当 i, j 指向的字符均为字母（或 '.' 可以看成一个特殊的字母）时，
//    只需判断对应位置的字符即可，
//    若相等，只需判断 i,j 之前的字符串是否匹配即可，转化为子问题 f[i-1][j-1].
//    若不等，则当前的 i,j 肯定不能匹配，为 false.
   
//        f[i-1][j-1]   i
//             |        |
//    S [a  b  b  b  b][c] 
   
//    P [a  b  *  d  *][c]
//                      |
//                      j
   

// 2. 如果当前 j 指向的字符为 '*'，则不妨把类似 'a*', 'b*' 等的当成整体看待。
//    看下面的例子

//             i
//             |
//    S  a  b [b] b  b  c  
   
//    P  a [b  *] d  *  c
//             |
//             j
   
//    注意到当 'b*' 匹配完 'b' 之后，它仍然可以继续发挥作用。
//    因此可以只把 i 前移一位，而不丢弃 'b*', 转化为子问题 f[i-1][j]:
   
//          i
//          | <--
//    S  a [b] b  b  b  c  
   
//    P  a [b  *] d  *  c
//             |
//             j
   
//    另外，也可以选择让 'b*' 不再进行匹配，把 'b*' 丢弃。
//    转化为子问题 f[i][j-2]:

//             i
//             |
//    S  a  b [b] b  b  c  
    
//    P [a] b  *  d  *  c
//       |
//       j <--
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if match(i, j) {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			if p[j-1] == '*' {
				if !match(i, j-1) {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}
// @lc code=end

