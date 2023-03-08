package leetcode

//给定二叉搜索树（BST）的根节点
// root 和一个整数值
// val。
//
// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回
// null 。
//
//
//
// 示例 1:
//
//
//
//
//
//输入：root = [4,2,7,1,3], val = 2
//输出：[2,1,3]
//
//
// 示例 2:
//
//
//输入：root = [4,2,7,1,3], val = 5
//输出：[]
//
//
//
//
// 提示：
//
//
// 数中节点数在 [1, 5000] 范围内
// 1 <= Node.val <= 10⁷
// root 是二叉搜索树
// 1 <= val <= 10⁷
//
//
// Related Topics 树 二叉搜索树 二叉树 👍 370 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//	type TreeNode struct {
//		    Val int
//		     Left *TreeNode
//		     Right *TreeNode
//		}
func searchBST(root *TreeNode, val int) *TreeNode {
	var help func(node *TreeNode)
	var t *TreeNode
	ok := false
	help = func(node *TreeNode) {
		if ok {
			return
		}
		if node == nil {
			return
		}
		help(node.Left)
		if node.Val == val {
			t = node
			ok = true
			return
		}
		help(node.Right)

	}
	help(root)
	return t

}

//leetcode submit region end(Prohibit modification and deletion)
