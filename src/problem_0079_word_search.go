// https://leetcode.cn/problems/word-search
// 单词搜索
package main

func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
			return false
		}
		if board[i][j] != word[k] {
			return false
		}
		board[i][j] ^= 255
		defer func() { board[i][j] ^= 255 }()
		return dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
	}

	for i, v := range board {
		for j, _ := range v {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
