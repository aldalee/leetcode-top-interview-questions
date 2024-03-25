// https://leetcode.cn/problems/increasing-triplet-subsequence/
// 递增的三元子序列
package main

import "sort"

func increasingTriplet(nums []int) bool {
	ends, right := make([]int, 3), 0
	ends[0] = nums[0]
	for _, num := range nums {
		left := sort.Search(right+1, func(i int) bool {
			return ends[i] >= num
		})
		if right = max(left, right); right > 1 {
			return true
		}
		ends[left] = num
	}
	return false
}
