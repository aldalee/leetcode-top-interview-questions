// https://leetcode.cn/problems/merge-intervals
// 合并区间
package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, interval := range intervals {
		L, R := interval[0], interval[1]
		n := len(merged)
		if n == 0 || merged[n-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			merged[n-1][1] = max(merged[n-1][1], R)
		}
	}
	return merged
}
