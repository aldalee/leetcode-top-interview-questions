// https://leetcode.cn/problems/two-sum/
// 两数之和
package main

func twoSum(nums []int, target int) []int {
	hash := map[int]int{} // <value, index>
	for idx, num := range nums {
		if v, ok := hash[target-num]; ok {
			return []int{v, idx}
		}
		hash[num] = idx
	}
	return []int{-1, -1}
}
