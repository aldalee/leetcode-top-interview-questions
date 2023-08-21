// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 电话号码的字母组合
package main

var phone = [][]rune{
	{'a', 'b', 'c'},      // 2
	{'d', 'e', 'f'},      // 3
	{'g', 'h', 'i'},      // 4
	{'j', 'k', 'l'},      // 5
	{'m', 'n', 'o'},      // 6
	{'p', 'q', 'r', 's'}, // 7
	{'t', 'u', 'v'},      // 8
	{'w', 'x', 'y', 'z'}, // 9
}

func letterCombinations(digits string) []string {
	var ans []string
	strs := []rune(digits)
	path := make([]rune, len(strs))
	f(strs, 0, path, &ans)
	return ans
}

func f(strs []rune, idx int, path []rune, ans *[]string) {
	if idx == len(strs) {
		*ans = append(*ans, string(path))
		return
	}
	cands := phone[strs[idx]-'2']
	for _, cur := range cands {
		path[idx] = cur
		f(strs, idx+1, path, ans)
	}
}
