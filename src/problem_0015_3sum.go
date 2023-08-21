// https://leetcode.cn/problems/3sum
// 三数之和
package main

import "sort"

func threeSum(nums []int) [][]int {
	return threeSums(nums, 0)
}

func threeSums(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if (3*nums[i] <= target) && (i == 0 || nums[i] != nums[i-1]) {
			res := twoSums(nums, i+1, target-nums[i])
			for _, cur := range res {
				cur = append(cur, nums[i])
				ans = append(ans, cur)
			}
		}
	}
	return ans
}

func twoSums(nums []int, begin, target int) [][]int {
	L, R := begin, len(nums)-1
	ans := make([][]int, 0)
	for L < R {
		sum := nums[L] + nums[R]
		if sum > target {
			R--
		} else if sum < target {
			L++
		} else {
			if L == begin || nums[L-1] != nums[L] {
				cur := []int{nums[L], nums[R]}
				ans = append(ans, cur)
			}
			L++
			R--
		}
	}
	return ans
}
