/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
import "strings"
func numJewelsInStones(J string, S string) int {
	if J == "" || S == "" {
		return 0
	}
	sum := 0
	for i := range J {
		sum += strings.Count(S,string(J[i]))
	}
	return sum
}
// @lc code=end

