// https://leetcode.cn/problems/trapping-rain-water
// 接雨水
package main

func trap(height []int) (water int) {
	n := len(height)
	L, R := 1, n-2
	maxL, maxR := height[0], height[n-1]
	for L <= R {
		if maxL < maxR {
			water += max(maxL-height[L], 0)
			maxL = max(maxL, height[L])
			L++
		} else {
			water += max(maxR-height[R], 0)
			maxR = max(maxR, height[R])
			R--
		}
	}
	return
}
