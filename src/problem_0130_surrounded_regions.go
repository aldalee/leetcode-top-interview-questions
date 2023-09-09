// https://leetcode.cn/problems/surrounded-regions
// 被围绕的区域
package main

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	var free func(i, j int)
	free = func(i, j int) {
		if i < 0 || i == n || j < 0 || j == m || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'F'
		free(i-1, j)
		free(i+1, j)
		free(i, j-1)
		free(i, j+1)
	}
	// in the vertical boundary direction
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			free(i, 0)
		}
		if board[i][m-1] == 'O' {
			free(i, m-1)
		}
	}
	// in the horizontal boundary direction
	for j := 1; j < m-1; j++ {
		if board[0][j] == 'O' {
			free(0, j)
		}
		if board[n-1][j] == 'O' {
			free(n-1, j)
		}
	}
	// modify
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}
}
