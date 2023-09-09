// https://leetcode.cn/problems/word-ladder/
// 单词接龙
package main

import "github.com/emirpasic/gods/sets/hashset"

func ladderLength(beginWord, endWord string, wordList []string) int {
	dict := hashset.New()
	for _, word := range wordList {
		dict.Add(word)
	}
	if !dict.Contains(endWord) {
		return 0
	}
	begin := hashset.New(beginWord)
	end := hashset.New(endWord)
	visited := hashset.New()
	for n := 2; !begin.Empty(); n++ {
		next := hashset.New()
		for _, word := range begin.Values() {
			for i := range word.(string) {
				w := []rune(word.(string))
				for c := 'a'; c <= 'z'; c++ {
					if c != w[i] {
						w[i] = c
						nextWord := string(w)
						if end.Contains(nextWord) {
							return n
						}
						if dict.Contains(nextWord) && !visited.Contains(nextWord) {
							next.Add(nextWord)
							visited.Add(nextWord)
						}
					}
				}
			}
		}
		begin = next
		if begin.Size() > end.Size() {
			begin, end = end, begin
		}
	}
	return 0
}
