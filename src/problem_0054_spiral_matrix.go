// https://leetcode.cn/problems/spiral-matrix
// 螺旋矩阵
package main

func spiralOrder(matrix [][]int) (ans []int) {
	x1, y1 := 0, 0
	x2, y2 := len(matrix)-1, len(matrix[0])-1
	for x1 <= x2 && y1 <= y2 {
		addEdge(matrix, x1, y1, x2, y2, &ans)
		x1++
		y1++
		x2--
		y2--
	}
	return
}

func addEdge(matrix [][]int, x1, y1, x2, y2 int, ans *[]int) {
	if x1 == x2 {
		for i := y1; i <= y2; i++ {
			*ans = append(*ans, matrix[x1][i])
		}
	} else if y1 == y2 {
		for i := x1; i <= x2; i++ {
			*ans = append(*ans, matrix[i][y1])
		}
	} else {
		row, col := x1, y1
		for ; col != y2; col++ {
			*ans = append(*ans, matrix[x1][col])
		}
		for ; row != x2; row++ {
			*ans = append(*ans, matrix[row][y2])
		}

		for ; col != y1; col-- {
			*ans = append(*ans, matrix[x2][col])
		}
		for ; row != x1; row-- {
			*ans = append(*ans, matrix[row][y1])
		}
	}
}
