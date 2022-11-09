package leetcode

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
//
// 示例 3：
//
//
//输入：root = []
//输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 1177 👎 0

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
//	Val int
//    Left *TreeNode
//    Right *TreeNode
//}

func isBalanced(root *TreeNode) bool {
	return Judge(root)

}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max1(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}

func Judge(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//	if root.Left==nil{}
	l := maxDepth1(root.Left)
	r := maxDepth1(root.Right)
	if l-r > 1 || r-l > 1 {
		return false
	}
	if !Judge(root.Left) {
		return false
	}
	if !Judge(root.Right) {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
