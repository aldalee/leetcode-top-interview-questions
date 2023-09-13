// https://leetcode.cn/problems/longest-palindromic-substring/
// 最长回文子串
// Manacher算法
package main

func longestPalindrome(s string) string {
	i, n := manacher(s)
	l := (i - n - 1) >> 1
	r := l + n
	return s[l:r]
}

func manacher(str string) (int, int) {
	s := change(str)
	n := len(s)
	p := make([]int, n)
	R, C := 0, 0
	idx, maxLen := 0, -1
	for i := 1; i < n-1; i++ {
		if i < R {
			p[i] = min(p[(C<<1)-i], R-i)
		} else {
			p[i] = 1
		}
		for s[i+p[i]] == s[i-p[i]] {
			p[i]++
		}
		if p[i]+i > R {
			R = p[i] + i
			C = i
		}
		if maxLen < p[i] {
			maxLen = p[i]
			idx = i
		}
	}
	return idx, maxLen - 1
}

func change(str string) []byte {
	n, k := len(str), 1
	s := make([]byte, (n<<1)+3)
	s[0] = '#'
	for i := 0; i < n; i++ {
		s[k] = '#'
		s[k+1] = str[i]
		k += 2
	}
	s[k] = '#'
	return s
}
