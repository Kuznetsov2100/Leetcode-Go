/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
import "strings"
func restoreIpAddresses(s string) []string {
	var res []string
	if len(s) < 4 || len(s) > 12 { // s长度小于4或大于12无法构成合法ip地址
		return res
	}
	var backtrack func(str string, ip []string)
	backtrack = func(str string, ip []string) {
		if len(ip) == 4 && str == "" { // ip字段为4个且当前可供选择字符串为空
			res = append(res, strings.Join(ip,"."))
			return
		}
		for i := 1; i <= 3; i++ {
			if i > len(str) { // 截取长度大于当前可供选择字符串的长度，退出
				return
			}
			if i == 3 && str[:i] > "255" { // 截取长度为3且截取的字符串大于“255”，退出
				return
			}
			if str[0] == '0' && i > 1 { // 截取长度大于1且含有前导0，退出
				return
			}
			ip = append(ip, str[:i])
			backtrack(str[i:],ip) // 递归剩下的可供选择的字符串
			ip = ip[:len(ip)-1]
		}
	}
	backtrack(s,[]string{})
	return res
}
// @lc code=end

