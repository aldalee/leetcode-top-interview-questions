// https://leetcode.cn/problems/rotate-image
// 旋转图像
package main

func rotate(matrix [][]int) {
	tr, tc := 0, 0
	dr, dc := len(matrix)-1, len(matrix)-1
	for tr < dr {
		rotateEdge(matrix, tr, tc, dr, dc)
		tr++
		tc++
		dr--
		dc--
	}
}

func rotateEdge(matrix [][]int, tr, tc, dr, dc int) {
	var tmp int
	for i := 0; i < dc-tc; i++ {
		tmp = matrix[tr][tc+i]
		matrix[tr][tc+i] = matrix[dr-i][tc] // 4 -> 1
		matrix[dr-i][tc] = matrix[dr][dc-i] // 3 -> 4
		matrix[dr][dc-i] = matrix[tr+i][dc] // 2 -> 3
		matrix[tr+i][dc] = tmp              // t -> 2
	}
}
