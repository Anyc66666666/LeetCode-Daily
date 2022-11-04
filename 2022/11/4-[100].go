package leetcode

import "fmt"

//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
//
//
// 示例 1：
//
//
//输入：p = [1,2,3], q = [1,2,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：p = [1,2], q = [1,null,2]
//输出：false
//
//
// 示例 3：
//
//
//输入：p = [1,2,1], q = [1,1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 两棵树上的节点数目都在范围 [0, 100] 内
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 923 👎 0

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
//   Val int
//   Left *TreeNode
//   Right *TreeNode
//}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var a, b []int
	var inorder, inorder1 func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		a = append(a, node.Val)
		inorder(node.Right)
	}
	inorder1 = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder1(node.Left)
		b = append(b, node.Val)
		inorder1(node.Right)
	}

	inorder(p)
	inorder1(q)
	if len(a) != len(b) {
		fmt.Println(len(a), len(b))
		return false
	}
	var l, r int
	for p != nil {
		l++
		p = p.Left
	}

	for q != nil {
		r++
		q = q.Left
	}
	if r != l {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Println(a[i], b[i])
			return false
		}
	}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)
