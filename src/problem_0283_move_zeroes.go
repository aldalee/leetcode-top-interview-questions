// https://leetcode.cn/problems/move-zeroes/
// 移动零
package main

func moveZeroes(nums []int) {
	for l, r, n := 0, 0, len(nums); r < n; r++ {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
