// https://leetcode.cn/problems/powx-n/
// Pow(x,n)
package main

func myPow(x float64, n int) float64 {
	pow := abs(n)
	t, ans := x, float64(1)
	for pow != 0 {
		if (pow & 1) != 0 {
			ans *= t
		}
		pow >>= 1
		t *= t
	}
	if n < 0 {
		return 1 / ans
	} else {
		return ans
	}
}
