// https://leetcode.cn/problems/word-search-ii/
// 单词搜索II
package main

type PrefixTree struct {
	pass  int
	end   int
	nexts [26]*PrefixTree
}

func (root *PrefixTree) Insert(word string) {
	node := root
	node.pass++
	for _, path := range word {
		path -= 'a'
		if node.nexts[path] == nil {
			node.nexts[path] = &PrefixTree{}
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

func findWords(board [][]byte, words []string) []string {
	root := &PrefixTree{}
	for _, word := range words {
		root.Insert(word)
	}
	var (
		path []rune
		res  []string
		dfs  func(int, int, *PrefixTree) int
		dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	)

	var toString = func(path []rune) string {
		str := make([]rune, len(path))
		for i, cha := range path {
			str[i] = cha
		}
		return string(str)
	}

	dfs = func(row, col int, cur *PrefixTree) int {
		ch := board[row][col]
		if ch == '#' {
			return 0
		}
		idx := int(ch - 'a')
		if cur.nexts[idx] == nil || cur.nexts[idx].pass == 0 {
			return 0
		}
		cur = cur.nexts[idx]
		path = append(path, rune(ch))
		fix := 0
		if cur.end > 0 {
			res = append(res, toString(path))
			cur.end--
			fix++
		}
		board[row][col] = '#'
		for _, dir := range dirs {
			r, c := row+dir[0], col+dir[1]
			if r >= 0 && c >= 0 && r < len(board) && c < len(board[0]) {
				fix += dfs(r, c, cur)
			}
		}
		board[row][col] = ch
		path = path[:len(path)-1]
		cur.pass -= fix
		return fix
	}

	for row, v := range board {
		for col, _ := range v {
			dfs(row, col, root)
		}
	}
	return res
}
