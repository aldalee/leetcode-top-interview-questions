// https://leetcode.cn/problems/roman-to-integer
// 罗马数字转整数
package main

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (res int) {
	n := len(s)
	for i := range s {
		v := m[s[i]]
		if i < n-1 && v < m[s[i+1]] {
			res -= v
		} else {
			res += v
		}
	}
	return
}
