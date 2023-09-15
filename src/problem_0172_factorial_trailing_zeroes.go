// https://leetcode.cn/problems/factorial-trailing-zeroes/
// 阶乘后的零
package main

func trailingZeroes(n int) (ans int) {
	for n != 0 {
		n /= 5
		ans += n
	}
	return
}
