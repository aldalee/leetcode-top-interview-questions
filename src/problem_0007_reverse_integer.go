// https://leetcode.cn/problems/reverse-integer
// 整数反转
package main

import "math"

func reverse(x int) (res int) {
	M := math.MaxInt32 / 10
	m := math.MinInt32 / 10
	for x != 0 {
		if res < m || res > M {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}
	return
}
