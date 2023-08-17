// https://leetcode.cn/problems/string-to-integer-atoi
// 字符串转换整数 (atoi)
package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = extractNumbers(strings.TrimSpace(s))
	str := []rune(s)
	if !valid(str) {
		return 0
	}
	return convertNumber(str)
}

func extractNumbers(s string) string {
	start, end := 0, len(s)
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		start = 1
	}
	for start < end && s[start] == '0' {
		start++
	}
	for i := len(s) - 1; i >= start; i-- {
		if s[i] < '0' || s[i] > '9' {
			end = i
		}
	}
	return s[0:start] + s[start:end]
}

func valid(s []rune) bool {
	if len(s) == 0 {
		return false
	}
	return (s[0] != '-' && s[0] != '+') || len(s) != 1
}

func convertNumber(s []rune) (res int) {
	positive := s[0] != '-'
	idx := 0
	if s[0] == '-' || s[0] == '+' {
		idx = 1
	}
	for idx < len(s) {
		res = res*10 + int(s[idx]-'0')
		if res > math.MaxInt32 {
			if positive {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		idx++
	}
	if positive {
		return res
	} else {
		return -res
	}
}
