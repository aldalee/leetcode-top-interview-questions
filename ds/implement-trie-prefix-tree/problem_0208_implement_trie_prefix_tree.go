// Package implement_trie_prefix_tree
// https://leetcode.cn/problems/implement-trie-prefix-tree/
// 实现Trie树（前缀树）
package implement_trie_prefix_tree

type Trie struct {
	nexts [26]*Trie
	end   bool
}

func Constructor() Trie {
	return Trie{}
}

func (root *Trie) Insert(word string) {
	node := root
	for _, path := range word {
		path -= 'a'
		if node.nexts[path] == nil {
			node.nexts[path] = &Trie{}
		}
		node = node.nexts[path]
	}
	node.end = true
}

func (root *Trie) Search(word string) bool {
	node := root.searchPrefix(word)
	return node != nil && node.end
}

func (root *Trie) StartsWith(prefix string) bool {
	return root.searchPrefix(prefix) != nil
}

func (root *Trie) searchPrefix(prefix string) *Trie {
	node := root
	for _, path := range prefix {
		path -= 'a'
		if node.nexts[path] == nil {
			return nil
		}
		node = node.nexts[path]
	}
	return node
}
