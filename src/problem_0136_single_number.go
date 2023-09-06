// https://leetcode.cn/problems/single-number/
// 只出现一次的数字
package main

func singleNumber(nums []int) (eor int) {
	for _, num := range nums {
		eor = eor ^ num
	}
	return
}
