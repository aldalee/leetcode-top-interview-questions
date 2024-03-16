// https://leetcode.cn/problems/odd-even-linked-list/
// 奇偶链表
package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, even, evenHead := head, head.Next, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
