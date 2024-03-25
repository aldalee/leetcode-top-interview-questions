// https://leetcode.cn/problems/longest-increasing-subsequence
// 最长递增子序列
package main

import "sort"

func lengthOfLIS(nums []int) int {
	var ends []int
	for _, num := range nums {
		left := sort.Search(len(ends), func(i int) bool {
			return ends[i] >= num
		})
		if left == len(ends) {
			ends = append(ends, num)
		} else {
			ends[left] = num
		}
	}
	return len(ends)
}
