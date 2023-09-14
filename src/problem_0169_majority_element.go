// https://leetcode.cn/problems/majority-element
// 多数元素
package main

func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			res = num
			cnt++
		} else if num == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
