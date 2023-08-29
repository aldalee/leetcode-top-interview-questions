// https://leetcode.cn/problems/sqrtx
// x的平方根
package main

func mySqrt(x int) (res int) {
	l, r, m := 0, x, 0
	for l <= r {
		m = ((r - l) >> 1) + l
		if m*m <= x {
			res = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return
}
