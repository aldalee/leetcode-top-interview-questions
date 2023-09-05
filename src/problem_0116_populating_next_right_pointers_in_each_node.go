// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
// 填充每个节点的下一个右侧节点指针
package main

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	for level := root; level.Left != nil; level = level.Left {
		for cur := level; cur != nil; cur = cur.Next {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
		}
	}
	return root
}
