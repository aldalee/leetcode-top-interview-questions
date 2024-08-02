// https://leetcode.cn/problems/missing-ranges/
// 缺失的区间
// Given a sorted integer array nums, where the range of elements
// are in the inclusive range [lower, upper], return its missing ranges.
// Example:
// Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
// Output: [[2, 2],[4, 49],[51, 74],[76, 99]]
package main

// TODO: 待验证
func findMissingRanges(nums []int, lower, upper int) (ans [][]int) {
	for _, num := range nums {
		if num > lower {
			ans = append(ans, []int{lower, num - 1})
		}
		lower = num + 1
	}
	if lower <= upper {
		ans = append(ans, []int{lower, upper})
	}
	return
}
