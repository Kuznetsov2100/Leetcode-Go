/*
 * @lc app=leetcode.cn id=1410 lang=golang
 *
 * [1410] HTML 实体解析器
 */

// @lc code=start
func entityParser(text string) string {
	return strings.NewReplacer("&quot;", "\"", "&apos;", "'", "&amp;", "&", "&gt;", ">", "&lt;", "<", "&frasl;", "/").Replace(text)
}
// @lc code=end

