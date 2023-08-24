// https://leetcode.cn/problems/first-missing-positive
// 缺失的第一个正数
package main

func firstMissingPositive(nums []int) int {
	// expect: nums[l] == l + 1
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == l+1 {
			l++
		} else if nums[l] <= l || nums[l] > r || nums[l] == nums[nums[l]-1] {
			r--
			swap(nums, l, r)
		} else {
			swap(nums, l, nums[l]-1)
		}
	}
	return l + 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
