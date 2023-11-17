package leetcode

//给定一个 n 叉树的根节点
// root ，返回 其节点值的 后序遍历 。
//
// n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：[5,6,3,2,4,1]
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]
//
//
//
//
// 提示：
//
//
// 节点总数在范围 [0, 10⁴] 内
// 0 <= Node.val <= 10⁴
// n 叉树的高度小于或等于 1000
//
//
//
//
// 进阶：递归法很简单，你可以使用迭代法完成此题吗?
//
// Related Topics 栈 树 深度优先搜索 👍 298 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//	type Node struct {
//		    Val int
//		    Children []*Node
//		}
func postorder(root *Node) []int {
	//var ans []int
	//var dfs func(node *Node)
	//dfs= func(node *Node) {
	//	if node==nil{
	//		return
	//	}
	//	for _,v:=range node.Children{
	//		dfs(v)
	//	}
	//	ans=append(ans,node.Val)
	//}
	//dfs(root)
	//return ans

	//------------

	var ans []int
	if root == nil {
		return ans
	}
	st := []*Node{root}
	vis := map[*Node]bool{}
	for len(st) > 0 {
		node := st[len(st)-1]
		// 如果当前节点为叶子节点或者当前节点的子节点已经遍历过
		if len(node.Children) == 0 || vis[node] {
			ans = append(ans, node.Val)
			st = st[:len(st)-1]
			continue
		}
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
		vis[node] = true
	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
