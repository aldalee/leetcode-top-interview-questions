// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
// 在排序数组中查找元素的第一个和最后一个位置
package main

import "sort"

func searchRange(nums []int, target int) []int {
	L := binarySearchLeftIdx(nums, target)
	R := binarySearchRightIdx(nums, target)
	return []int{L, R}
}

func binarySearchLeftIdx(arr []int, num int) int {
	l, r := 0, len(arr)-1
	idx := -1
	for l <= r {
		m := l + ((r - l) >> 1)
		if arr[m] < num {
			l = m + 1
		} else if arr[m] > num {
			r = m - 1
		} else {
			idx = m
			r = m - 1
		}
	}
	return idx
}

func binarySearchRightIdx(arr []int, num int) int {
	l, r := 0, len(arr)-1
	idx := -1
	for l <= r {
		m := l + ((r - l) >> 1)
		if arr[m] < num {
			l = m + 1
		} else if arr[m] > num {
			r = m - 1
		} else {
			idx = m
			l = m + 1
		}
	}
	return idx
}

// SearchRange 利用标准库
func SearchRange(nums []int, target int) []int {
	L := sort.SearchInts(nums, target)
	if L == len(nums) || nums[L] != target {
		return []int{-1, -1}
	}
	R := sort.SearchInts(nums, target+1) - 1
	return []int{L, R}
}
