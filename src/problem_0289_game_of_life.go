// https://leetcode.cn/problems/game-of-life/
// 生命游戏
package main

func gameOfLife(board [][]int) {
	one := func(i, j int) int {
		if i >= 0 && i < len(board) && j >= 0 &&
			j < len(board[0]) && (board[i][j]&1 == 1) {
			return 1
		}
		return 0
	}
	// number of '1' around board[i][j]
	neighbors := func(i, j int) int {
		return one(i-1, j-1) + one(i-1, j) + one(i-1, j+1) + one(i, j-1) +
			one(i, j+1) + one(i+1, j-1) + one(i+1, j) + one(i+1, j+1)
	}
	// set next generation's life state
	for i, vector := range board {
		for j, life := range vector {
			ones := neighbors(i, j)
			if ones == 3 || (life == 1 && ones == 2) {
				board[i][j] |= 2
			}
		}
	}
	// get next generation's life state
	for i, vector := range board {
		for j, _ := range vector {
			board[i][j] >>= 1
		}
	}
}
