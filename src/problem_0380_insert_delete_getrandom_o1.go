// https://leetcode.cn/problems/insert-delete-getrandom-o1/
// O(1) 时间插入、删除和获取随机元素
package main

import "math/rand/v2"

type RandomizedSet struct {
	valMap map[int]int
	IdxMap map[int]int
	size   int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valMap: make(map[int]int),
		IdxMap: make(map[int]int),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.valMap[val]; ok {
		return false
	}
	s.valMap[val] = s.size
	s.IdxMap[s.size] = val
	s.size++
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.valMap[val]; !ok {
		return false
	}
	delIdx, lastIdx := s.valMap[val], s.size-1
	lastVal := s.IdxMap[lastIdx]
	s.valMap[lastVal] = delIdx
	s.IdxMap[delIdx] = lastVal
	delete(s.valMap, val)
	delete(s.IdxMap, lastIdx)
	s.size--
	return true
}

func (s *RandomizedSet) GetRandom() int {
	return s.IdxMap[rand.IntN(s.size)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
