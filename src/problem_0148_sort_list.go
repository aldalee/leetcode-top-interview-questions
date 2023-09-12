// https://leetcode.cn/problems/sort-list
// 排序链表
package main

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		dummy = &ListNode{0, head}
		prev  *ListNode
		curr  = head
		size  = 0
	)
	for ; curr != nil; size++ {
		curr = curr.Next
	}
	for step := 1; step < size; step <<= 1 {
		prev = dummy
		curr = dummy.Next
		for curr != nil {
			head1 := curr
			for i := 1; i < step && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			curr.Next = nil
			curr = head2
			for i := 1; i < step && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}
			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}
			prev.Next = mergeTwoLists(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			curr = next
		}
	}
	return dummy.Next
}
