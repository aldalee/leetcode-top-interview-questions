// https://leetcode.cn/problems/perfect-squares/
// 完全平方数
package main

import "math"

func numSquares(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}
	for i := 1; i*i < n; i++ {
		if isPerfectSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isPerfectSquare(a int) bool {
	b := int(math.Sqrt(float64(a)))
	return b*b == a
}
