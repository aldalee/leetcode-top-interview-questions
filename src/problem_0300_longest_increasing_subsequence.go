// https://leetcode.cn/problems/longest-increasing-subsequence
// 最长递增子序列
package main

import "sort"

func lengthOfLIS(nums []int) int {
	var ends []int
	for i := 0; i < len(nums); i++ {
		// find the first position that greater than or equal num in ends
		idx := sort.Search(len(ends), func(j int) bool {
			return ends[j] >= nums[i]
		})
		if idx == len(ends) {
			ends = append(ends, nums[i])
		} else {
			ends[idx] = nums[i]
		}
	}
	return len(ends)
}
