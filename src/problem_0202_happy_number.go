// https://leetcode.cn/problems/happy-number/
// 快乐数
package main

func isHappy(n int) bool {
	var step = func(n int) (sum int) {
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return
	}
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}
