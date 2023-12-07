package leetcode

//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [2,1,3]
//输出: 1
//
//
// 示例 2:
//
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 547 👎 0

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
//		    Left *TreeNode
//		    Right *TreeNode
//		}
func findBottomLeftValue(root *TreeNode) int {
	//queue:=list.New()
	//var gradation int
	//queue.PushBack(root)
	//for queue.Len()>0{
	//	length:=queue.Len()
	//	for i:=0;i<length;i++{
	//		node:=queue.Remove(queue.Front()).(*TreeNode)
	//		if i==0{gradation=node.Val}
	//		if node.Left!=nil{
	//			queue.PushBack(node.Left)
	//		}
	//		if node.Right!=nil{
	//			queue.PushBack(node.Right)
	//		}
	//	}
	//}
	//return gradation
	depth := 0
	var dfs func(node *TreeNode, depth int)
	var a [][]int
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(a) == depth {
			a = append(a, []int{})
		}
		a[depth] = append(a[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, depth)
	return a[len(a)-1][0]
}

//leetcode submit region end(Prohibit modification and deletion)
