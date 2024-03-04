// https://leetcode.cn/problems/delete-node-in-a-linked-list/
// 删除链表中的节点
package main

func deleteNode(node *ListNode) {
	// In fact, it does not delete the actual node, but the node.next,
	// which appears to have deleted the node.
	// So this question is not very interesting.
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
