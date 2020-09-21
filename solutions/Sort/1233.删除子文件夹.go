/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 */

// @lc code=start
func removeSubfolders(folder []string) []string {
	// 先对folder进行字典序排序，则具有相同前缀的文件夹一定靠在一起
	// 遍历folder，假如f没有前缀prefix，说明这个f非子文件夹，将其加入结果数组，
	// 同是更新prefix
	var res []string
	sort.Strings(folder)
	prefix := "#"
	for _, f := range folder {
		if !strings.HasPrefix(f, prefix) {
			res = append(res, f)
			prefix = f + "/"
		}
	}
	return res
}
// @lc code=end

