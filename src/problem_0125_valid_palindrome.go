// https://leetcode.cn/problems/valid-palindrome/
// 验证回文串
package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !isLetterOrDigit(rune(s[l])) {
			l++
		} else if !isLetterOrDigit(rune(s[r])) {
			r--
		} else if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
