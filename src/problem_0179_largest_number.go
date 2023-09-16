// https://leetcode.cn/problems/largest-number/
// 最大数
package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[j]+strs[i] < strs[i]+strs[j]
	})
	var ans strings.Builder
	for _, str := range strs {
		ans.WriteString(str)
	}
	if ans.String()[0] == '0' {
		return "0"
	}
	return ans.String()
}
