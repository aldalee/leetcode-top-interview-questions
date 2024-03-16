// https://leetcode.cn/problems/odd-even-linked-list/
// 奇偶链表
package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy, n = &ListNode{}, 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
		if cur.Next == nil && n%2 == 0 {
			cur.Next = dummy
			break
		}
	}
	var pre = modifyList(head, dummy)
	if n%2 == 0 {
		pre.Next = dummy.Next
	}
	return head
}

func modifyList(head, dummy *ListNode) (pre *ListNode) {
	var cur, evenHead = head, head.Next
	for next := head; next.Next != nil; {
		next = next.Next
		if next.Next == dummy {
			pre = cur
		}
		cur.Next = next.Next
		cur = next
	}
	cur.Next = evenHead
	return pre
}
