// https://leetcode.cn/problems/search-in-rotated-sorted-array/
// 搜索旋转排序数组
package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := ((r - l) >> 1) + l
		if nums[m] == target {
			return m
		}
		if (nums[m] >= nums[0]) == (target >= nums[0]) {
			if nums[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] >= nums[0] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
