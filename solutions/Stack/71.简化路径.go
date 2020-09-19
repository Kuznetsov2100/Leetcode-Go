/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	var stack []string
	for _, s := range strings.Split(path, "/") {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 出栈
			}
		} else if s != "" && s != "." { // s非空且s != "."时，入栈
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}
// @lc code=end

