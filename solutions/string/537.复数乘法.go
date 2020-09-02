/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 */

// @lc code=start
func complexNumberMultiply(a string, b string) string {
	num1plusindex := strings.Index(a, "+")
	num2plusindex := strings.Index(b, "+")
	num1real, _ := strconv.Atoi(a[:num1plusindex])
	num1imag, _ := strconv.Atoi(a[num1plusindex+1:len(a)-1])
	num2real, _ := strconv.Atoi(b[:num2plusindex])
	num2imag, _ := strconv.Atoi(b[num2plusindex+1:len(b)-1])
	var res strings.Builder
	res.WriteString(strconv.Itoa(num1real*num2real - num1imag*num2imag))
	res.WriteString("+")
	res.WriteString(strconv.Itoa(num1real*num2imag + num1imag*num2real))
	res.WriteString("i")
	return res.String()
}
// @lc code=end

