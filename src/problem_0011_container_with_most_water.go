// https://leetcode.cn/problems/container-with-most-water
// 盛最多水的容器
package main

func maxArea(h []int) (res int) {
	l, r := 0, len(h)-1
	for l < r {
		area := min(h[l], h[r]) * (r - l)
		res = max(res, area)
		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}
	return
}
