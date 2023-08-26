// https://leetcode.cn/problems/jump-game
// 跳跃游戏
package main

func canJump(nums []int) bool {
	maxJump := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+nums[i])
	}
	return true
}
