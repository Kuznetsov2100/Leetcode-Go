/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
import "strconv"
import "bytes"
func calculate(s string) int {
	var M,Q []string
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	var operand bytes.Buffer
	for i := 0; i < len(s); i++ {	
		if s[i] >= '0' && s[i] <= '9' {	
			for ;i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				operand.WriteByte(s[i])
			}
			Q = append(Q, operand.String())
			operand.Reset()
			i--
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			for len(M) > 0 && precedence[M[len(M)-1]] >= precedence[string(s[i])]  {
				Q = append(Q, M[len(M)-1])
				M = M[:len(M)-1]		
			}
			M = append(M,string(s[i]))
		} else {
			continue
		}
	}
	if len(M) == 0 {
		if num, ok := isNumber(Q[len(Q)-1]); ok {
			return num
		}
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

