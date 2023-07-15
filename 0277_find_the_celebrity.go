// https://leetcode.cn/problems/find-the-celebrity/
// 搜寻名人
package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		ans := 0
		for i := 1; i < n; i++ {
			if knows(ans, i) {
				ans = i
			}
		}
		for i := 0; i < n; i++ {
			if ans == i {
				continue
			}
			if knows(ans, i) || !knows(i, ans) {
				return -1
			}
		}
		return ans
	}
}
