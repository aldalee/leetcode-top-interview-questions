// https://leetcode.cn/problems/copy-list-with-random-pointer/
// 复制带随机指针的链表
package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val    int
 *     Next   *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var next, clone *Node
	for cur := head; cur != nil; cur = next {
		next = cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
	}
	for cur := head; cur != nil; cur = next {
		next = cur.Next.Next
		clone = cur.Next
		if cur.Random != nil {
			clone.Random = cur.Random.Next
		}
	}
	var res = head.Next
	for cur := head; cur != nil; cur = next {
		next = cur.Next.Next
		clone = cur.Next
		cur.Next = next
		if next != nil {
			clone.Next = next.Next
		}
	}
	return res
}
