// https://leetcode.cn/problems/search-a-2d-matrix-ii/
// 搜索二维矩阵 II
package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for x, y := 0, n-1; x < m && y >= 0; {
		if matrix[x][y] < target {
			x++
		} else if matrix[x][y] > target {
			y--
		} else {
			return true
		}
	}
	return false
}
