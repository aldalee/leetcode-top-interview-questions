// https://leetcode.cn/problems/evaluate-reverse-polish-notation
// 逆波兰表达式求值
package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	var compute = func(op string) {
		n := len(stack)
		y, x := stack[n-1], stack[n-2]
		stack = stack[:n-2]
		switch op {
		case "+":
			stack = append(stack, x+y)
		case "-":
			stack = append(stack, x-y)
		case "*":
			stack = append(stack, x*y)
		case "/":
			stack = append(stack, x/y)
		}
	}
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			compute(token)
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}
