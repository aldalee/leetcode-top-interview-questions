// https://leetcode.cn/problems/sudoku-solver
// 解数独
package main

var (
	row  = make([][]bool, 9)
	col  = make([][]bool, 9)
	grid = make([][]bool, 9)
)

func solveSudoku(board [][]byte) {
	pre(board)
	dfs(board, 0, 0)
}

func pre(board [][]byte) {
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10)
		col[i] = make([]bool, 10)
		grid[i] = make([]bool, 10)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			k := 3*(i/3) + j/3
			if board[i][j] != '.' {
				x := board[i][j] - '0'
				row[i][x] = true
				col[j][x] = true
				grid[k][x] = true
			}
		}
	}
}

func dfs(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	ni, nj := i, j+1
	if j == 8 {
		ni = i + 1
		nj = 0
	}
	if board[i][j] != '.' {
		return dfs(board, ni, nj)
	} else {
		k := 3*(i/3) + j/3
		for x := 1; x <= 9; x++ {
			if (!row[i][x]) && (!col[j][x]) && (!grid[k][x]) {
				fill(board, i, j, k, x)
				if dfs(board, ni, nj) {
					return true
				}
				restore(board, i, j, k, x)
			}
		}
		return false
	}
}

func fill(board [][]byte, i, j, k, x int) {
	row[i][x] = true
	col[j][x] = true
	grid[k][x] = true
	board[i][j] = byte(x + '0')
}

func restore(board [][]byte, i, j, k, x int) {
	row[i][x] = false
	col[j][x] = false
	grid[k][x] = false
	board[i][j] = '.'
}
