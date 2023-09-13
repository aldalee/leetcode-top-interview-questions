// https://leetcode.cn/problems/maximum-product-subarray/
// 乘积最大子数组
package main

func maxProduct(nums []int) int {
	ans := nums[0]
	preMax, preMin := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMin := min(nums[i], preMin*nums[i], preMax*nums[i])
		curMax := max(nums[i], preMin*nums[i], preMax*nums[i])
		preMin, preMax = curMin, curMax
		ans = max(ans, preMax)
	}
	return ans
}
