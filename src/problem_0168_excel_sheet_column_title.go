// https://leetcode.cn/problems/excel-sheet-column-title/
// Excel表列名称
package main

import "strings"

func convertToTitle(columnNumber int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var ans strings.Builder
	var reverse = func(s string) string {
		r := []rune(s)
		for i, n := 0, len(s); i < n/2; i++ {
			r[i], r[n-i-1] = r[n-i-1], r[i]
		}
		return string(r)
	}
	for columnNumber > 0 {
		columnNumber--
		ans.WriteByte(alphabet[columnNumber%26])
		columnNumber /= 26
	}
	return reverse(ans.String())
}
