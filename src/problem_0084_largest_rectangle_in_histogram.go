// https://leetcode.cn/problems/largest-rectangle-in-histogram
// 柱状图中最大的矩形
package main

func largestRectangleArea(heights []int) (maxArea int) {
	stack := make([]int, 0) // monotonic-stack store the index
	for i := 0; i <= len(heights); i++ {
		for len(stack) > 0 &&
			(i == len(heights) || heights[i] <= heights[stack[len(stack)-1]]) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k := -1
			if len(stack) > 0 {
				k = stack[len(stack)-1]
			}
			curArea := (i - k - 1) * heights[j]
			if curArea > maxArea {
				maxArea = curArea
			}
		}
		stack = append(stack, i)
	}
	return
}
