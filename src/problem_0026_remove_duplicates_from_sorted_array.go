// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// 删除有序数组中的重复项
package main

func removeDuplicates(nums []int) int {
	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[p] {
			p++
			nums[p] = nums[i]
		}
	}
	return p + 1
}
