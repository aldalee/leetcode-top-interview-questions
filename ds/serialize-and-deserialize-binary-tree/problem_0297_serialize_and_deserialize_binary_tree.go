// Package serialize_and_deserialize_binary_tree
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
// 二叉树的序列化与反序列化
// 腾讯课堂进阶班LeetCode 22节 开头视频
package serialize_and_deserialize_binary_tree

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return "TODO"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
