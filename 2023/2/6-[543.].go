package leetcode

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
//
//
// 示例 : 给定二叉树
//
//           1
//         / \
//        2   3
//       / \
//      4   5
//
//
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。
//
// Related Topics 树 深度优先搜索 二叉树 👍 1231 👎 0

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
//	    Val int
//	    Left *TreeNode
//	    Right *TreeNode
//}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil {
		return maxDia
	}
	var help func(node *TreeNode) int
	help = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := help(node.Left)
		rh := help(node.Right)
		maxDia = max(maxDia, lh+rh)
		return max(lh, rh) + 1

	}

	help(root)
	return maxDia

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

//leetcode submit region end(Prohibit modification and deletion)
