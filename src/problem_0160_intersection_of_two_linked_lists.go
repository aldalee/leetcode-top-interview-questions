// https://leetcode.cn/problems/intersection-of-two-linked-lists
// 相交链表
package main

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	cur1, cur2 := head1, head2
	for cur1 != cur2 {
		if cur1 != nil {
			cur1 = cur1.Next
		} else {
			cur1 = head2
		}
		if cur2 != nil {
			cur2 = cur2.Next
		} else {
			cur2 = head1
		}
	}
	return cur1
}
