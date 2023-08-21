// https://leetcode.cn/problems/divide-two-integers/
// 两数相除
package main

import "math"

// 位运算实现除法
func divide(dividend int, divisor int) int {
	// 除数是系统最小值
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	// 除数不是系统最小值，被除数是系统最小值
	if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		}
		r := div(add(dividend, 1), divisor)
		d := minus(dividend, multi(r, divisor))
		return add(r, div(d, divisor))
	}
	// 被除数不是系统最小值，除数不是系统最小值
	return div(dividend, divisor)
}

// 位运算实现加法
func add(a int, b int) int {
	s := a
	for b != 0 {
		s = a ^ b        // 无进位相加的信息
		b = (a & b) << 1 // 进位的信息
		a = s
	}
	return s
}

// 位运算实现减法
func minus(a int, b int) int {
	return add(a, negate(b))
}

// 位运算实现乘法
func multi(a int, b int) (r int) {
	for b != 0 {
		if b&1 != 0 {
			r = add(r, a)
		}
		a <<= 1
		b = int(uint(b) >> 1) // b >>>= 1
	}
	return
}

// 位运算实现正数除法
func div(a int, b int) (r int) {
	x, y := abs(a), abs(b)
	for i := 31; i > negate(1); i = minus(i, 1) {
		if x>>i >= y {
			r |= 1 << i
			x = minus(x, multi(y, 1<<i))
		}
	}
	if isNeg(a) != isNeg(b) {
		return negate(r)
	}
	return r
}

// 相反数
func negate(n int) int {
	return add(^n, 1)
}

// 判断负数
func isNeg(n int) bool {
	return n>>31 == -1
}

func abs(n int) int {
	if isNeg(n) {
		return negate(n)
	}
	return n
}

func printBinary(n int) string {
	res := make([]byte, 0)
	for i := 31; i >= 0; i-- {
		bit := (n >> i) & 1
		res = append(res, byte('0'+bit))
	}
	return string(res)
}
