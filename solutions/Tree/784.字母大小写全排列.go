/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
func letterCasePermutation(S string) []string {
	// dfs递归，遍历每一个位置的字符；
	// 保持当前位置的字符不变，递归下一个位置；
	// 如果当前位置为字母，则再加一个大小写转换之后的递归分支。
	var res []string
	var dfs func(char []byte, i int)
	dfs = func(char []byte, i int) {
		if i == len(S) {
			res = append(res, string(char))
			return
		}
		dfs(char,i+1)
		if char[i] >= 65 && char[i] <= 90 {
			char[i] += 32
			dfs(char,i+1)
		}else if char[i] >= 97 && char[i] <= 122 {
			char[i] -= 32
			dfs(char, i+1)
		}
	}
	dfs([]byte(S),0)
	return res
}
// @lc code=end

