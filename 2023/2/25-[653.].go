package leetcode

//给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
//
//
// 示例 1：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
//
//
// 示例 2：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1, 10⁴].
// -10⁴ <= Node.val <= 10⁴
// 题目数据保证，输入的 root 是一棵 有效 的二叉搜索树
// -10⁵ <= k <= 10⁵
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 哈希表 双指针 二叉树 👍 454 👎 0

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
//	    Right *TreeNode
//	 }
func findTarget(root *TreeNode, k int) bool {
	var help func(node *TreeNode)
	var a []int
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		help(node.Left)
		a = append(a, node.Val)
		help(node.Right)
	}

	help(root)

	if len(a) < 2 {
		return false
	}

	//1 2 4 5 6 8

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == k {
				return true
			}
		}
	}

	return false

}

//leetcode submit region end(Prohibit modification and deletion)
/*
func findTarget(root *TreeNode, k int) bool {
    set := map[int]struct{}{}
    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        if _, ok := set[k-node.Val]; ok {
            return true
        }
        set[node.Val] = struct{}{}
        return dfs(node.Left) || dfs(node.Right)
    }
    return dfs(root)
}

我们可以使用深度优先搜索的方式遍历整棵树，用哈希表记录遍历过的节点的值。

对于一个值为 x 的节点，我们检查哈希表中是否存在k−x 即可。如果存在对应的元素，那么我们就可以在该树上找到两个节点的和为 k；否则，我们将 x 放入到哈希表中。

如果遍历完整棵树都不存在对应的元素，那么该树上不存在两个和为 kkk 的节点。


*/
