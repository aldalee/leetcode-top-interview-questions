// https://leetcode.cn/problems/count-primes/
// 计数质数
package main

func countPrimes(n int) (cnt int) {
	sieve := make([]bool, n)
	for i := 2; i < n; i++ {
		if !sieve[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				sieve[j] = true
			}
		}
	}
	return
}
