package leetcode

//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//
// 如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
//
// 我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
//
// 只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,4], x = 4, y = 3
//输出：false
//
//
// 示例 2：
//
//
//输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
//输出：true
//
//
// 示例 3：
//
//
//
//
//输入：root = [1,2,3,null,4], x = 2, y = 3
//输出：false
//
//
//
// 提示：
//
//
// 二叉树的节点数介于 2 到 100 之间。
// 每个节点的值都是唯一的、范围为 1 到 100 的整数。
//
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 296 👎 0

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
//	     Val int
//	     Left *TreeNode
//	     Right *TreeNode
//	}

func isCousins(root *TreeNode, x int, y int) bool {
	//要想判断两个节点 xxx 和 yyy 是否为堂兄弟节点，
	//我们就需要求出这两个节点分别的「深度」以及「父节点」。
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var help func(node, Parent *TreeNode, depth int)
	help = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		if xFound && yFound {
			return
		}

		help(node.Left, node, depth+1)

		if xFound && yFound {
			return
		}

		help(node.Right, node, depth+1)

	}

	help(root, nil, 0)

	return xDepth == yDepth && xParent != yParent

}

//leetcode submit region end(Prohibit modification and deletion)
