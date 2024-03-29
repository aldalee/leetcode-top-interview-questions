// https://leetcode.cn/problems/missing-number/
// 丢失的数字
package main

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
