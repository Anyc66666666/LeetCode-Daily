package leetcode

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2177 👎 0

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
//    Val int
//    Left *TreeNode
//   Right *TreeNode
// } //

func isSymmetric(root *TreeNode) bool {
	var inorder func(n, m *TreeNode)
	var flag int
	inorder = func(n, m *TreeNode) {
		if (n == nil && m != nil) || (m == nil && n != nil) {
			flag = 6
			return
		}
		if n == nil && m == nil {
			return
		}
		if n.Val != m.Val {
			flag = 6

		}

		inorder(n.Left, m.Right)
		inorder(n.Right, m.Left)

	}
	a, b := root, root
	inorder(a, b)
	return flag != 6

}

//leetcode submit region end(Prohibit modification and deletion)
