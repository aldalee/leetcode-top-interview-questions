// https://leetcode.cn/problems/reverse-bits/
// 颠倒二进制位
package main

func reverseBits(n uint32) uint32 {
	masks := []uint32{
		0x0000ffff, // 00000000000000001111111111111111
		0x00ff00ff, // 00000000111111110000000011111111
		0x0f0f0f0f, // 00001111000011110000111100001111
		0x33333333, // 00110011001100110011001100110011
		0x55555555, // 01010101010101010101010101010101
	}
	shifts := []uint32{16, 8, 4, 2, 1}
	for i, mask := range masks {
		n = n>>shifts[i]&mask | n&mask<<shifts[i]
	}
	return n
}

//func reverseBits(n uint32) (rev uint32) {
//	for i := 31; i >= 0 && n > 0; i-- {
//		rev |= (n & 1) << i
//		n >>= 1
//	}
//	return
//}
