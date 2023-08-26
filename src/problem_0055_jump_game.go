// https://leetcode.cn/problems/jump-game
// 跳跃游戏
package main

func canJump(nums []int) bool {
	maxJump := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		if maxJump > len(nums) {
			return true
		}
		maxJump = max(maxJump, i+nums[i])
	}
	return true
}
