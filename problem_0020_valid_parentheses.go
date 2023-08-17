// https://leetcode.cn/problems/valid-parentheses/
// 有效的括号
package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			n := len(stack)
			if n == 0 {
				return false
			}
			top := stack[n-1]
			if (top == '(' && c == ')') ||
				(top == '[' && c == ']') ||
				(top == '{' && c == '}') {
				stack = stack[:n-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
