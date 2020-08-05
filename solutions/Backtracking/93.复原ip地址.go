/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
import "strings"
func restoreIpAddresses(s string) []string {
	var res []string
	length := len(s)
	if length < 4 || length > 12 { // s长度小于4或大于12无法构成合法ip地址
		return res
	}
	var backtrack func(pos int, substrlen int, ip []string)
	backtrack = func(pos int, substrlen int, ip []string) {
		if len(ip) == 4 && substrlen == 0 { // ip字段为4个且当前可供选择字符串为空
			res = append(res, strings.Join(ip,"."))
			return
		}
		for i := 1; i <= 3; i++ {
			if i > substrlen { // 截取长度大于当前可供选择字符串的长度，退出
				return
			}
			if i == 3 && s[pos:pos+i] > "255" { // 截取长度为3且截取的字符串大于“255”，退出
				return
			}
			if s[pos:pos+1] == "0" && i > 1 { // 截取长度大于1且含有前导0，退出
				return
			}
			ip = append(ip, s[pos:pos+i])
			backtrack(pos+i,substrlen-i,ip) // 递归剩下的可供选择的字符串
			ip = ip[:len(ip)-1]
		}
	}
	backtrack(0, length, []string{})
	return res
}
// @lc code=end

