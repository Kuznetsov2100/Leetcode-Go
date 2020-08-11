/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	var stack []int
	for len(tokens) > 0 {
		element := tokens[0]
		tokens = tokens[1:]
		if element != "+" && element != "-" && element != "*" && element != "/" {
			num, _ := strconv.Atoi(element)
			stack = append(stack, num)			
		} else {
			if element == "+" {
				numA := stack[len(stack)-1]
				numB := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, numB+numA)
			} else if element == "-" {
				numA := stack[len(stack)-1]
				numB := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, numB-numA)
			} else if element == "*" {
				numA := stack[len(stack)-1]
				numB := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, numB*numA)
			} else if element == "/" {
				numA := stack[len(stack)-1]
				numB := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, numB/numA)
			}
		}
	}
	return stack[0]
}
// @lc code=end

