// https://leetcode.cn/problems/fraction-to-recurring-decimal
// 分数到小数
package main

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	var ans []byte
	var abs = func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	if numerator != 0 && (numerator > 0) != (denominator > 0) {
		ans = append(ans, '-')
	}
	a, b := abs(int64(numerator)), abs(int64(denominator))
	ans = strconv.AppendInt(ans, a/b, 10)
	a %= b
	if a == 0 {
		return string(ans)
	}
	ans = append(ans, '.')
	m := map[int64]int{}
	for a != 0 {
		if i, ok := m[a]; ok {
			ans = append(ans[:i], append(append([]byte{'('}, ans[i:]...), ')')...)
			break
		}
		m[a] = len(ans)
		a *= 10
		ans = strconv.AppendInt(ans, a/b, 10)
		a %= b
	}
	return string(ans)
}
