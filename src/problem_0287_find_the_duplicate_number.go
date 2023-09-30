// https://leetcode.cn/problems/find-the-duplicate-number/
// 寻找重复数
package main

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}
