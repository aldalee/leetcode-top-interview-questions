// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/
// 矩阵中的最长递增路径
package main

func longestIncreasingPath(matrix [][]int) (ans int) {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i, vector := range matrix {
		for j := range vector {
			ans = max(LIP(matrix, i, j, dp), ans)
		}
	}
	return
}

func LIP(m [][]int, i, j int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	var top, down, left, right int
	if i > 0 && m[i][j] > m[i-1][j] {
		top = LIP(m, i-1, j, dp)
	}
	if i < len(m)-1 && m[i][j] > m[i+1][j] {
		down = LIP(m, i+1, j, dp)
	}
	if j > 0 && m[i][j] > m[i][j-1] {
		left = LIP(m, i, j-1, dp)
	}
	if j < len(m[0])-1 && m[i][j] > m[i][j+1] {
		right = LIP(m, i, j+1, dp)
	}
	dp[i][j] = max(top, down, left, right) + 1
	return dp[i][j]
}
