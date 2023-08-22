// https://leetcode.cn/problems/valid-sudoku
// 有效的数独
package main

func isValidSudoku(board [][]byte) bool {
	row := make([]int, 9)
	col := make([]int, 9)
	grid := make([]int, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			k := 3*(i/3) + j/3
			if board[i][j] != '.' {
				x := 1 << (board[i][j] - '0')
				if (row[i]&x) != 0 || (col[j]&x) != 0 || (grid[k]&x) != 0 {
					return false
				}
				row[i] |= x
				col[j] |= x
				grid[k] |= x
			}
		}
	}
	return true
}
