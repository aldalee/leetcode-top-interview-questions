// https://leetcode.cn/problems/permutations/
// 全排列
package main

func permute(nums []int) (ans [][]int) {
	generate(nums, 0, &ans)
	return
}

func generate(nums []int, i int, ans *[][]int) {
	if i == len(nums) {
		cur := make([]int, len(nums))
		copy(cur, nums)
		*ans = append(*ans, cur)
	} else {
		for j := i; j < len(nums); j++ {
			swap(nums, i, j)
			generate(nums, i+1, ans)
			swap(nums, i, j) // restore context
		}
	}
}
