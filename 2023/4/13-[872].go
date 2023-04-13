package leetcode

//请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
//
//
//
// 举个例子，如上图所示，给定一棵叶值序列为 (6, 7, 4, 9, 8) 的树。
//
// 如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
//
// 如果给定的两个根结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。
//
//
//
// 示例 1：
//
//
//
//
//输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,
//null,null,null,null,9,8]
//输出：true
//
//
// 示例 2：
//
//
//
//
//输入：root1 = [1,2,3], root2 = [1,3,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 给定的两棵树结点数在 [1, 200] 范围内
// 给定的两棵树上的值在 [0, 200] 范围内
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 204 👎 0

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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var a, b []int
	var help, help1 func(node *TreeNode)
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		help(node.Left)
		if node.Left == nil && node.Right == nil {
			a = append(a, node.Val)
		}
		help(node.Right)
	}
	help1 = func(node *TreeNode) {
		if node == nil {
			return
		}
		help1(node.Left)
		if node.Left == nil && node.Right == nil {
			b = append(b, node.Val)
		}
		help1(node.Right)
	}
	help(root1)
	help1(root2)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
