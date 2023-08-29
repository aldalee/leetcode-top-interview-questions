// https://leetcode.cn/problems/unique-paths
// 不同路径
package main

import "math/big"

// C(m+n-2,m-1)
func uniquePaths(m int, n int) int {
	N := int64(n + m - 2)
	M := int64(m - 1)
	cnt := new(big.Int).Binomial(N, M).Int64()
	return int(cnt)
}
