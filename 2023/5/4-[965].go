package leetcode

//如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
//
// 只有给定的树是单值二叉树时，才返回 true；否则返回 false。
//
//
//
// 示例 1：
//
//
//
// 输入：[1,1,1,1,1,null,1]
//输出：true
//
//
// 示例 2：
//
//
//
// 输入：[2,2,2,5,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 给定树的节点数范围是 [1, 100]。
// 每个节点的值都是整数，范围为 [0, 99] 。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 180 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	m := make(map[int]struct{})
	var ok bool
	var help func(node *TreeNode)
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		if ok {
			return
		}

		help(node.Left)
		m[node.Val] = struct{}{}
		if len(m) > 1 {
			ok = true
		}
		help(node.Right)

	}
	help(root)
	return !ok

}

//leetcode submit region end(Prohibit modification and deletion)
