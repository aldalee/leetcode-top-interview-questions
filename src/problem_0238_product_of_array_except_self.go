// https://leetcode.cn/problems/product-of-array-except-self/
// 除自身以外数组的乘积
package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = nums[i-1] * ans[i-1] // L[i] = nums[i-1] * L[i-1]
	}
	R := 1 // R = R[i]
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * R // ans[i] = L[i] * R[i]
		R = nums[i] * R     // R[i+1] = nums[i] * R[i]
	}
	return ans
}
