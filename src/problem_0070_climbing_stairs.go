// https://leetcode.cn/problems/climbing-stairs/
// 爬楼梯
package main

func climbStairs(n int) int {
	M := [][]int{{1, 1}, {1, 0}}
	res := matrixPower(M, n)
	return res[0][0]
}

func matrixPower(matrix [][]int, n int) [][]int {
	res := make([][]int, len(matrix))
	for i := range res {
		res[i] = make([]int, len(res))
		res[i][i] = 1
	}
	t := matrix
	for ; n > 0; n >>= 1 {
		if (n & 1) != 0 {
			res = multiply(res, t)
		}
		t = multiply(t, t)
	}
	return res
}

// A[n*k] B[k*m]
func multiply(A, B [][]int) [][]int {
	n, k, m := len(A), len(A[0]), len(B[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for c := 0; c < k; c++ {
				res[i][j] += A[i][c] * B[c][j]
			}
		}
	}
	return res
}
