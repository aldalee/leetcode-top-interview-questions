// https://leetcode.cn/problems/find-peak-element/
// 寻找峰值
package main

func findPeakElement(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	l, r := 0, n-1
	for l < r {
		m := ((r - l) >> 1) + l
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
