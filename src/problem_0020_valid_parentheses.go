// https://leetcode.cn/problems/valid-parentheses/
// 有效的括号
package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := range s {
		if m[s[i]] > 0 {
			n := len(stack)
			if n == 0 || stack[n-1] != m[s[i]] {
				return false
			}
			stack = stack[:n-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
