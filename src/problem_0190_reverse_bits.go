// https://leetcode.cn/problems/reverse-bits/
// 颠倒二进制位
package main

func reverseBits(n uint32) uint32 {
	masks := []uint32{
		0x0000ffff, // binary: 16 zeros,16 ones
		0x00ff00ff, // binary:  8 zeros, 8 ones ...
		0x0f0f0f0f, // binary:  4 zeros, 4 ones ...
		0x33333333, // binary:  2 zeros, 2 ones ...
		0x55555555, // binary:  1 zeros, 1 ones ...
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
