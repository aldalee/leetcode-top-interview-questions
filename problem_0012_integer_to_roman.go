// https://leetcode.cn/problems/integer-to-roman/
// 整数转罗马数字
package main

import "strings"

var c = [][]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	{"", "M", "MM", "MMM"},
}
var d = [4]int{1, 10, 100, 1000}

func intToRoman(num int) string {
	var roman strings.Builder
	for i := 3; i >= 0; i-- {
		roman.WriteString(c[i][num/d[i]%10])
	}
	return roman.String()
}
