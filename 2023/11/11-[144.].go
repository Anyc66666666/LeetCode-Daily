package leetcode

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1173 👎 0

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
//   Val int
//   Left *TreeNode
//   Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	//var a []int
	//var dfs func(node *TreeNode)
	//dfs= func(node *TreeNode) {
	//	if node==nil{
	//		return
	//	}
	//
	//	a=append(a,node.Val)
	//
	//	dfs(node.Left)
	//	dfs(node.Right)
	//}
	//dfs(root)
	//return a

	//--------------------------
	//var ans []int
	//if root==nil{
	//	return ans
	//}
	//st:=list.New()
	//st.PushBack(root)
	//for st.Len()>0{
	//	node:=st.Remove(st.Back()).(*TreeNode)
	//	ans=append(ans,node.Val)
	//	if node.Right!=nil{
	//		st.PushBack(node.Right)
	//	}
	//	if node.Left!=nil{
	//		st.PushBack(node.Left)
	//	}
	//}
	//return ans

	///---------------
	var vals []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return vals
}

//leetcode submit region end(Prohibit modification and deletion)
