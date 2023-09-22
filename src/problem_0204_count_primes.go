// https://leetcode.cn/problems/count-primes/
// 计数质数
package main

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	sieve := make([]bool, n)
	count := n / 2 // exclude all even numbers.
	for i := 3; i*i < n; i += 2 {
		if !sieve[i] {
			for j := i * i; j < n; j += 2 * i {
				if !sieve[j] {
					count--
					sieve[j] = true
				}
			}
		}
	}
	return count
}
