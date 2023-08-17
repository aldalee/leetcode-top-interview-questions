// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
// 删除链表的倒数第N个结点
package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	cur, pre := head, dummy
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
