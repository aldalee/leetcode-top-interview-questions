// https://leetcode.cn/problems/two-sum/
// 两数之和
package main

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for idx, num := range nums {
		if val, ok := hash[target-num]; ok {
			return []int{val, idx}
		}
		hash[num] = idx
	}
	return []int{-1, -1}
}
