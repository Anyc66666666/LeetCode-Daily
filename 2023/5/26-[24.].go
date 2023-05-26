package leetcode

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
// Related Topics 递归 链表 👍 1830 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//   Val int
//   Next *ListNode
//}

func swapPairs(head *ListNode) *ListNode {

	dummyHead := &ListNode{0, head}
	temp := dummyHead

	// 1 2 3 4
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next      // 1     ->2
		node2 := temp.Next.Next // 2   ->3 ..

		temp.Next = node2       // [] -> 2
		node1.Next = node2.Next // 1  -> 3 ..
		node2.Next = node1      // [] -> 2 -> 1 -> 3 ..

		temp = temp.Next.Next

		// 2 1 3 4

	}

	return dummyHead.Next

}

//leetcode submit region end(Prohibit modification and deletion)
