// https://leetcode.cn/problems/number-of-islands/
// 岛屿数量
package main

func numIslands(grid [][]byte) (islands int) {
	row, col := len(grid), len(grid[0])
	var connect func(i, j int)
	connect = func(i, j int) {
		if i < 0 || i == row || j < 0 || j == col || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		connect(i-1, j)
		connect(i+1, j)
		connect(i, j-1)
		connect(i, j+1)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				islands++
				connect(i, j)
			}
		}
	}
	return
}
