// https://leetcode.cn/problems/power-of-three/
// 3的幂
package main

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
