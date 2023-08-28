// https://leetcode.cn/problems/house-robber
// 打家劫舍
package main

func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// 空间压缩
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	pre2Sum := nums[0]
	pre1Sum := max(nums[0], nums[1])
	curSum := pre1Sum
	for i := 2; i < n; i++ {
		curSum = max(pre1Sum, nums[i]+pre2Sum)
		pre2Sum, pre1Sum = pre1Sum, curSum
	}
	return curSum
}
