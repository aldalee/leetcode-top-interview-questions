// https://leetcode.cn/problems/jump-game-ii
// 跳跃游戏 II
package main

func jump(nums []int) (step int) {
	curJump, maxJump := 0, nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if curJump < idx {
			step++
			curJump = maxJump
		}
		maxJump = max(maxJump, idx+nums[idx])
	}
	return
}
