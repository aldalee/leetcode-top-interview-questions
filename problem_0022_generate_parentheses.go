// https://leetcode.cn/problems/generate-parentheses
// 括号生成
package main

func generateParenthesis(n int) []string {
	var ans []string
	path := make([]byte, n<<1)
	DFS(path, 0, 0, n, &ans)
	return ans
}

// DFS 深度优先遍历，依次在path上作决定（边尝试边剪枝）
//
// Params:
//
//	path – path[0...idx-1]已经填写完成
//	idx  – 当前要作决定的位置
//	num  – 左括号-右括号的数量（约束参数）
//	rest – 左括号剩余的数量（约束参数）
//	ans –  存储答案
func DFS(path []byte, idx, num, rest int, ans *[]string) {
	if idx == len(path) {
		*ans = append(*ans, string(path))
	} else {
		if rest > 0 {
			path[idx] = '('
			DFS(path, idx+1, num+1, rest-1, ans)
		}
		if num > 0 {
			path[idx] = ')'
			DFS(path, idx+1, num-1, rest, ans)
		}
	}
}
