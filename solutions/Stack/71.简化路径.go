/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	var res strings.Builder
	var stack []string
	str := strings.Split(path, "/")
	for _, s := range str {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 出栈
			}
		} else if s != "" && s != "." { // s非空且s != "."时，入栈
			stack = append(stack, s)
		}
	}
	if len(stack) == 0 { // 栈为空， 返回"/"
		return "/"
	}
	// 从栈底到栈顶，拼接结果字符串
	for _, s := range stack { 
		res.WriteString("/"+s)
	}
	return res.String()
}
// @lc code=end

