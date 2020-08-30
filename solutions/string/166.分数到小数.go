/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var res string
	if (numerator ^ denominator) < 0 { // 判断numerator 与 denominator是否异号
		res += "-"
	}
	numerator, denominator = abs(numerator), abs(denominator) // numerator 与 denominator都取绝对值
	res += strconv.Itoa(numerator/denominator) // 先放入两数相除的整数部分
	remainder := numerator % denominator
	if remainder == 0 { // 整除,直接返回
		return res
	}
	res += "."
	m := make(map[int]int) // 记录余数出现的位置
	dup_pos := -1
	// 模拟竖式除法
	for remainder != 0 {
		if pos, ok := m[remainder]; ok {
			dup_pos = pos
			break
		}
		m[remainder] = len(res)
		remainder *= 10
		res += strconv.Itoa(remainder/denominator)
		remainder %= denominator
	}
	if dup_pos == -1 { // 余数不循环
		return res
	}
	return fmt.Sprintf("%s(%s)", res[:dup_pos], res[dup_pos:])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

