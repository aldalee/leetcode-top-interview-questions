// https://leetcode.cn/problems/contains-duplicate/
// 存在重复元素
package main

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
