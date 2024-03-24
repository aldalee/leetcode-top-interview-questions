// https://leetcode.cn/problems/4sum-ii/
// 四数相加 II
package main

func fourSumCount(a, b, c, d []int) (ans int) {
	m := make(map[int]int, len(a)*len(b))
	for _, u := range a {
		for _, v := range b {
			m[u+v]++
		}
	}
	for _, u := range c {
		for _, v := range d {
			ans += m[-u-v]
		}
	}
	return
}
