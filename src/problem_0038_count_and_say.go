// https://leetcode.cn/problems/count-and-say
// 外观数列
package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for l, r := 0, 0; r < len(str); l = r {
			for r < len(str) && str[l] == str[r] {
				r++
			}
			cur.WriteString(strconv.Itoa(r - l))
			cur.WriteByte(str[l])
		}
		str = cur.String()
	}
	return str
}
