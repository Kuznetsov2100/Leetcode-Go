/*
 * @lc app=leetcode.cn id=1447 lang=golang
 *
 * [1447] 最简分数
 */

// @lc code=start
func simplifiedFractions(n int) []string {
	var res []string
	for i := 1; i < n; i++ {
		for j := i+1; j <= n; j++ {
			if gcd(i, j) == 1  {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return res
}


func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}
// @lc code=end

