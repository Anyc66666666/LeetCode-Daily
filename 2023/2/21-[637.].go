package leetcode

//给定一个非空二叉树的根节点
// root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10⁻⁵ 以内的答案可以被接受。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[3.00000,14.50000,11.00000]
//解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
//因此返回 [3, 14.5, 11] 。
//
//
// 示例 2:
//
//
//
//
//输入：root = [3,9,20,15,7]
//输出：[3.00000,14.50000,11.00000]
//
//
//
//
// 提示：
//
//
//
//
//
// 树中节点数量在 [1, 10⁴] 范围内
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 394 👎 0

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
// Val int
// Left *TreeNode
// Right *TreeNode
//}
type data struct {
	sum   int //每一层的总和
	count int //每一层的节点数
}

func averageOfLevels(root *TreeNode) []float64 {
	var levelData []data //索引相当于第几层(level)
	var help func(node *TreeNode, level int)
	help = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level < len(levelData) {
			levelData[level].sum += node.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{node.Val, 1})
		}

		help(node.Left, level+1)
		help(node.Right, level+1)

	}
	help(root, 0)

	averages := make([]float64, len(levelData))
	for i, d := range levelData {
		averages[i] = float64(d.sum) / float64(d.count)
	}
	return averages

}

//leetcode submit region end(Prohibit modification and deletion)

/*
使用深度优先搜索计算二叉树的层平均值，需要维护两个数组，
counts 用于存储二叉树的每一层的节点数，sums用于存储二叉树的每一层的节点值之和。
搜索过程中需要记录当前节点所在层，如果访问到的节点在第 i层，
则将 counts[i] 的值加 1，并将该节点的值加到sums[i]。

遍历结束之后，第 i 层的平均值即为 sums[i]/counts[i]


*/
