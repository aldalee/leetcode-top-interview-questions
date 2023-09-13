// https://leetcode.cn/problems/evaluate-reverse-polish-notation
// 逆波兰表达式求值
package main

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := linkedliststack.New()
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			compute(stack, token)
		} else {
			stack.Push(num)
		}
	}
	v, _ := stack.Pop()
	return v.(int)
}

func compute(stack *linkedliststack.Stack, op string) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	ans := 0
	switch op {
	case "+":
		ans = x.(int) + y.(int)
	case "-":
		ans = x.(int) - y.(int)
	case "*":
		ans = x.(int) * y.(int)
	case "/":
		ans = x.(int) / y.(int)
	}
	stack.Push(ans)
}
