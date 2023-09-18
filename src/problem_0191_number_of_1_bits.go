// https://leetcode.cn/problems/number-of-1-bits/
// 位1的个数
package main

func hammingWeight(n uint32) (bits int) {
	for bits = 0; n != 0; bits++ {
		n &= n - 1
	}
	return
}
