// https://leetcode.cn/problems/longest-increasing-subsequence
// 最长递增子序列
package main

func lengthOfLIS(nums []int) int {
	var ends []int
	// find the first position that >= num
	var binarySearchLeftIdx = func(num int) int {
		idx := -1
		for l, r := 0, len(ends)-1; l <= r; {
			m := l + ((r - l) >> 1)
			if ends[m] >= num {
				idx = m
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return idx
	}
	for i := 0; i < len(nums); i++ {
		idx := binarySearchLeftIdx(nums[i])
		if idx == -1 {
			ends = append(ends, nums[i])
		} else {
			ends[idx] = nums[i]
		}
	}
	return len(ends)
}
