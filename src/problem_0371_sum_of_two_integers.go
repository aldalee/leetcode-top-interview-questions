// https://leetcode.cn/problems/sum-of-two-integers/
// 两整数之和
package main

func getSum(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}
