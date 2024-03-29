// https://leetcode.cn/problems/valid-palindrome-ii/
// 验证回文串II
package main

func validPalindrome(s string) bool {
	var isPalindrome = func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		l++
		r--
	}
	return true
}
