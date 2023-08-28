// https://leetcode.cn/problems/maximum-subarray
// 最大子数组和
package main

func maxSubArray(nums []int) int {
	pre := nums[0]
	ans := pre
	for i := 1; i < len(nums); i++ {
		pre = max(nums[i], nums[i]+pre)
		ans = max(ans, pre)
	}
	return ans
}
