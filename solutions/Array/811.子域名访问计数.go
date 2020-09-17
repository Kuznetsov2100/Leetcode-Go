/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	var res []string
	m := make(map[string]int)
	for _, domain := range cpdomains {
		parts := strings.Split(domain, " ")
		count, _ := strconv.Atoi(parts[0])
		firstdotindex := strings.IndexByte(parts[1], '.')
		seconddotindex := strings.LastIndexByte(parts[1], '.')
		if firstdotindex != seconddotindex {
			m[parts[1][seconddotindex+1:]] += count
		}
		m[parts[1]] += count
		m[parts[1][firstdotindex+1:]] += count
	}
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+ " " + k)
	}
	return res
}
// @lc code=end

