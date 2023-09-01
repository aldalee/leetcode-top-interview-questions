// https://leetcode.cn/problems/sort-colors/
// 颜色分类
package main

func sortColors(nums []int) {
	less := -1        // 小于划分值区域右边界
	more := len(nums) // 大于划分值区域左边界
	pivot := 1        // 划分值
	for idx := 0; idx < more; {
		if nums[idx] == pivot {
			idx++
		} else if nums[idx] < pivot {
			less++
			swap(nums, idx, less)
			idx++
		} else {
			more--
			swap(nums, idx, more)
		}
	}
}
