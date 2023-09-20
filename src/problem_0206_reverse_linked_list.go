// https://leetcode.cn/problems/reverse-linked-list/
// 反转链表
package main

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
