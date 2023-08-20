package leetcode

import "math"

//给定一个二叉树，我们在树的节点上安装摄像头。
//
// 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
//
// 计算监控树的所有节点所需的最小摄像头数量。
//
//
//
// 示例 1：
//
//
//
// 输入：[0,0,null,0,0]
//输出：1
//解释：如图所示，一台摄像头足以监控所有节点。
//
//
// 示例 2：
//
//
//
// 输入：[0,0,null,0,null,0,null,null,0]
//输出：2
//解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。
//
//
// 提示：
//
//
// 给定树的节点数的范围是 [1, 1000]。
// 每个节点的值都是 0。
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 638 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
  type TreeNode struct {
   Val int
   Left *TreeNode
     Right *TreeNode
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	const inf = math.MaxInt32 / 2

	var dfs func(node *TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = min3(a, min3(la+rb, ra+lb))
		c = min3(a, lb+rb)
		return
	}

	_, ans, _ := dfs(root)
	return ans

}
func min3(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
