// https://leetcode.cn/problems/number-of-1-bits/
// 位1的个数
package main

func hammingWeight(n uint32) int {
	masks := []uint32{0x55555555, 0x33333333, 0x0f0f0f0f, 0x00ff00ff, 0x0000ffff}
	shifts := []uint32{1, 2, 4, 8, 16}
	for i, mask := range masks {
		n = n&mask + n>>shifts[i]&mask
	}
	return int(n)
}

//func hammingWeight(n uint32) (bits int) {
//	for bits = 0; n != 0; bits++ {
//		n &= n - 1
//	}
//	return
//}
