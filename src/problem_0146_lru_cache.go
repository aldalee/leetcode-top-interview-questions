// https://leetcode.cn/problems/lru-cache/
// LRU 缓存
package main

type LRUCache struct {
	capacity int
	kmap     map[int]*DNode
	list     *DoublyLinkedList
}

type DNode struct {
	k, v       int
	prev, next *DNode
}

type DoublyLinkedList struct {
	head *DNode
	tail *DNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		make(map[int]*DNode),
		NewDoublyLinkedList(),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.kmap[key]; ok {
		this.list.moveToTail(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.kmap[key]; ok {
		node.v = value
		this.list.moveToTail(node)
	} else {
		if len(this.kmap) == this.capacity {
			this.removeMostUnusedCache()
		}
		node := &DNode{k: key, v: value}
		this.kmap[key] = node
		this.list.addToTail(node)
	}
}

func (this *LRUCache) removeMostUnusedCache() {
	node := this.list.removeHead()
	delete(this.kmap, node.k)
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &DNode{}
	tail := &DNode{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head, tail}
}

func (l *DoublyLinkedList) addToTail(node *DNode) {
	if node == nil {
		return
	}
	node.prev = l.tail.prev
	node.next = l.tail
	l.tail.prev.next = node
	l.tail.prev = node
}

func (l *DoublyLinkedList) moveToTail(node *DNode) {
	l.removeDLinkedNode(node)
	l.addToTail(node)
}

func (l *DoublyLinkedList) removeHead() *DNode {
	res := l.head.next
	if res == l.tail {
		return nil
	}
	l.removeDLinkedNode(res)
	return res
}

func (l *DoublyLinkedList) removeDLinkedNode(node *DNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
