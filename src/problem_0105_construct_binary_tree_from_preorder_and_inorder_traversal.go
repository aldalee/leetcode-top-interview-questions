// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 从前序与中序遍历序列构造二叉树
package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := map[int]int{}
	for i, v := range inorder {
		m[v] = i
	}

	var f func(l1, r1, l2, r2 int) *TreeNode
	f = func(l1, r1, l2, r2 int) *TreeNode {
		if l1 > r1 {
			return nil
		}
		root := &TreeNode{Val: preorder[l1]}
		if l1 == r1 {
			return root
		}
		h := m[preorder[l1]]
		root.Left = f(l1+1, l1+h-l2, l2, h-1)
		root.Right = f(l1+h-l2+1, r1, h+1, r2)
		return root
	}
	return f(0, len(preorder)-1, 0, len(inorder)-1)
}
