/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	var stack []int
	for i := range tokens {
		element := tokens[i]
		if element != "+" && element != "-" && element != "*" && element != "/" {
			num, _ := strconv.Atoi(element)
			stack = append(stack, num)			
		} else {
			numA := stack[len(stack)-1]
			numB := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if element == "+" {	
				stack = append(stack, numB+numA)
			} else if element == "-" {
				stack = append(stack, numB-numA)
			} else if element == "*" {
				stack = append(stack, numB*numA)
			} else if element == "/" {
				stack = append(stack, numB/numA)
			}
		}
	}
	return stack[0]
}
// @lc code=end

