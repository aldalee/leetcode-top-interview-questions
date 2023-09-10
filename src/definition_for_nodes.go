package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node Definition for a Node.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Next   *Node
	Random *Node
}
