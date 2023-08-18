// https://leetcode.cn/problems/merge-two-sorted-lists/
// 合并两个有序链表
package main

func mergeTwoLists(L1, L2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for L1 != nil && L2 != nil {
		if L1.Val <= L2.Val {
			pre.Next = L1
			L1 = L1.Next
		} else {
			pre.Next = L2
			L2 = L2.Next
		}
		pre = pre.Next
	}
	if L1 == nil {
		pre.Next = L2
	} else {
		pre.Next = L1
	}
	return dummy.Next
}
