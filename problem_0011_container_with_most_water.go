// https://leetcode.cn/problems/container-with-most-water
// 盛最多水的容器
package main

import "math"

func maxArea(h []int) (res int) {
	l, r := 0, len(h)-1
	for l < r {
		area := (int)(math.Min(float64(h[l]), float64(h[r]))) * (r - l)
		res = (int)(math.Max(float64(res), float64(area)))
		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}
	return
}
