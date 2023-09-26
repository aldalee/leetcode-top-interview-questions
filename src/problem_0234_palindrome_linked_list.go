// https://leetcode.cn/problems/palindrome-linked-list/
// 回文链表
package main

func isPalindrome(head *ListNode) bool {
	if head != nil {
		return true
	}
	middleUpper := func() *ListNode {
		fast, slow := head, head
		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}()
	middleLower := reverseList(middleUpper.Next)
	p, q := head, middleLower
	res := true
	for q != nil {
		if p.Val != q.Val {
			res = false
			break
		}
		p = p.Next
		q = q.Next
	}
	middleUpper.Next = reverseList(middleLower)
	return res
}
