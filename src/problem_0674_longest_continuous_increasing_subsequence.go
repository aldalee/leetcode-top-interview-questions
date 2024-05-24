// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
// 最长连续递增序列
package main

func findLengthOfLCIS(nums []int) (ans int) {
	l, r := 0, 1
	for l < len(nums) {
		for r < len(nums) && nums[r] > nums[r-1] {
			r++
		}
		ans = max(ans, r-l)
		l = r
		r++
	}
	return
}
