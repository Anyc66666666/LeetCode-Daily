package leetcode

import (
	"strconv"
)

//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 852 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
//}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	helper(root, "")
	return paths
}

func helper(node *TreeNode, path string) {
	if node != nil {
		p := path
		p = p + strconv.Itoa(node.Val)

		if node.Right == nil && node.Left == nil {
			paths = append(paths, p)
		} else {
			p = p + "->"
			helper(node.Left, p)
			helper(node.Right, p)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
