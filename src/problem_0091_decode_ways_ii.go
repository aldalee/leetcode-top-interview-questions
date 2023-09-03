// https://leetcode.cn/problems/decode-ways-ii/
// 解码方法 II
package main

func NumDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		switch s[i] {
		case '0':
		case '*':
			dp[i] = dp[i+1] * 9
			if i+1 < n {
				switch s[i+1] {
				case '*':
					dp[i] += dp[i+2] * 15
				default:
					if s[i+1] < '7' {
						dp[i] += dp[i+2] * 2
					} else {
						dp[i] += dp[i+2]
					}
				}
			}
		default:
			dp[i] = dp[i+1]
			if i+1 < n {
				switch s[i+1] {
				case '*':
					if s[i] < '3' {
						if s[i] == '1' {
							dp[i] += dp[i+2] * 9
						} else {
							dp[i] += dp[i+2] * 6
						}
					}
				default:
					if (s[i]-'0')*10+(s[i+1]-'0') < 27 {
						dp[i] += dp[i+2]
					}
				}
			}
		}
		dp[i] %= 1000000007
	}
	return dp[0]
}

// 递归函数
func decode(s string, idx int) (ways int) {
	if idx == len(s) {
		return 1
	}
	if s[idx] == '0' {
		return 0
	}
	switch s[idx] {
	case '*':
		ways = decode(s, idx+1) * 9
		if idx+1 == len(s) {
			return ways
		}
		switch s[idx+1] {
		case '*':
			ways += decode(s, idx+2) * 15
		default:
			if s[idx+1] < '7' {
				ways += decode(s, idx+2) * 2
			} else {
				ways += decode(s, idx+2)
			}
		}
	default:
		ways = decode(s, idx+1)
		if idx+1 == len(s) {
			return ways
		}
		switch s[idx+1] {
		case '*':
			if s[idx] < '3' {
				if s[idx] == '1' {
					ways += decode(s, idx+2) * 9
				} else {
					ways += decode(s, idx+2) * 6
				}
			}
		default:
			if (s[idx]-'0')*10+(s[idx+1]-'0') < 27 {
				ways += decode(s, idx+2)
			}
		}
	}
	return
}
