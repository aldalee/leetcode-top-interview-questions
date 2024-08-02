// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
// 有序矩阵中第K小的元素
package main

import "math"

func kthSmallest2(matrix [][]int, k int) (ans int) {
	n := len(matrix)
	noMoreNum := func(value int) (int, int) {
		near, num := math.MinInt32, 0
		row, col := 0, n-1
		for row < n && col >= 0 {
			if matrix[row][col] <= value {
				near = max(near, matrix[row][col])
				num += col + 1
				row++
			} else {
				col--
			}
		}
		return near, num
	}

	minVal, maxVal := matrix[0][0], matrix[n-1][n-1]
	for minVal <= maxVal {
		midVal := minVal + ((maxVal - minVal) >> 1)
		near, cnt := noMoreNum(midVal)
		if cnt < k {
			minVal = midVal + 1
		} else {
			ans = near
			maxVal = midVal - 1
		}
	}
	return
}
