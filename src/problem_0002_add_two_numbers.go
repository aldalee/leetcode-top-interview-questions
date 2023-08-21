// https://leetcode.cn/problems/add-two-numbers/
// 两数相加
package main

func addTwoNumbers(L1, L2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for L1 != nil || L2 != nil {
		x, y := 0, 0
		if L1 != nil {
			x = L1.Val
			L1 = L1.Next
		}
		if L2 != nil {
			y = L2.Val
			L2 = L2.Next
		}
		sum := x + y + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}
