/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
import "strconv"
func calculate(s string) int {
	var M,Q []string
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	operand := 0
	pushOperand := func() {
		if operand != 0 {
			Q = append(Q, strconv.Itoa(operand))
			operand = 0
		}
	}
	for i := range s {	
		if s[i] >= '0' && s[i] <= '9' {
			if i == 0 && s[i] == '0' {
				Q = append(Q, "0")
			} else if i > 0 && s[i] == '0' && (s[i-1] < '0' || s[i-1] > '9') {
				Q = append(Q, "0")
			} else {
				operand = 10 * operand + int(s[i]- '0')
			}
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			pushOperand()
			for len(M) > 0 && precedence[M[len(M)-1]] >= precedence[string(s[i])]  {
				Q = append(Q, M[len(M)-1])
				M = M[:len(M)-1]		
			}
			M = append(M,string(s[i]))
		} 
		if i == len(s)-1 {	
			pushOperand()
		}
	}
	if num, ok := isNumber(Q[len(Q)-1]); len(M) == 0 && ok {
		return num
	}
	for len(M) > 0 {
		Q = append(Q, M[len(M)-1])
		M = M[:len(M)-1]
	}
	return evalRPN(Q)
}

func isNumber(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	return num, err == nil
}

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

