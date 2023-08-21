// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
// 找出字符串中第一个匹配项的下标
// Knuth-Morris-Pratt算法
// RabinKarp算法
package main

func strStr(haystack string, needle string) int {
	if len(haystack) < 1 || len(haystack) < len(needle) {
		return -1
	}
	return KMP(haystack, needle)
	//return RabinKarp(haystack,needle)
	//return strings.Index(haystack, needle)
}

func KMP(T, P string) int {
	x, y := 0, 0
	next := getNext(P)
	for x < len(T) && y < len(P) {
		if T[x] == P[y] {
			x++
			y++
		} else if next[y] == -1 { // y = 0
			x++
		} else {
			y = next[y]
		}
	}
	idx := x - y
	if y != len(P) {
		idx = -1
	}
	return idx
}

// 前缀串与后缀串的最长匹配长度
func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	idx, pn := 2, 0
	for idx < len(next) {
		if pattern[idx-1] == pattern[pn] {
			pn++
			next[idx] = pn
			idx++
		} else if pn > 0 {
			pn = next[pn]
		} else {
			next[idx] = 0
			idx++
		}
	}
	return next
}

const PrimeRK = 16777619

func RabinKarp(T, P string) int {
	n, m := len(T), len(P)
	ht, hp, pow := HashStr(T, P)
	for i := 0; i < n-m+1; i++ {
		if ht == hp && T[i:i+m] == P {
			return i
		}
		if i < n-m {
			ht = ht*PrimeRK + uint32(T[i+m]) - pow*uint32(T[i])
		}
	}
	return -1
}

// HashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func HashStr(T, P string) (uint32, uint32, uint32) {
	ht, hp := uint32(0), uint32(0)
	for i := 0; i < len(P); i++ {
		hp = hp*PrimeRK + uint32(P[i])
		ht = ht*PrimeRK + uint32(T[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(P); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return ht, hp, pow
}
