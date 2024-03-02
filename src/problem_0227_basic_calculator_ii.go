// https://leetcode.cn/problems/basic-calculator-ii/
// 基本计算器 II
package main

import "unicode"

func calculate(s string) (ans int) {
	var stack []int
	preSign, num := '+', 0
	for i, ch := range s {
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		}
		if !unicode.IsDigit(ch) && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign, num = ch, 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}
