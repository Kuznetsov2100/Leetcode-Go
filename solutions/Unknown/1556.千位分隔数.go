/*
 * @lc app=leetcode.cn id=1556 lang=golang
 *
 * [1556] 千位分隔数
 */

// @lc code=start
func thousandSeparator(n int) string {
	strn := strconv.Itoa(n)
	if n < 1000 {
		return strn
	}
	var res []byte
	cnt := 1
	for i := len(strn)-1; i >= 0; i-- {
		res = append(res, strn[i])
		if cnt % 3 == 0 && i != 0 {
			res = append(res, '.')
		}
		cnt++
	}

	// reverse res
	for i , j := 0, len(res)-1; i < j; i,j=i+1,j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
// @lc code=end

