package leetcode

//给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
//
//
//
// 示例 1：
//
//
//输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
//输出：32
//
//
// 示例 2：
//
//
//输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
//输出：23
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 2 * 10⁴] 内
// 1 <= Node.val <= 10⁵
// 1 <= low <= high <= 10⁵
// 所有 Node.val 互不相同
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 320 👎 0

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
//	    Val int
//	    Left *TreeNode
//	    Right *TreeNode
//	}
func rangeSumBST(root *TreeNode, low int, high int) int {
	var help func(node *TreeNode)
	var sum int
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		help(node.Left)
		if node.Val >= low && node.Val <= high {
			sum = sum + node.Val
		}
		help(node.Right)
	}

	help(root)
	return sum

}

//leetcode submit region end(Prohibit modification and deletion)
