// https://leetcode.cn/problems/happy-number/
// 快乐数
package main

func isHappy(n int) bool {
	var next = func(n int) (sum int) {
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return
	}
	slow, fast := n, next(n)
	for fast != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}
