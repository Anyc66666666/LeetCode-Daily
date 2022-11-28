package leetcode

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics 递归 链表 👍 2859 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	     Val int
//	     Next *ListNode
//	 }
func reverseList(head *ListNode) *ListNode {
	var a []int
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	m := new(ListNode)
	m.Val = 0
	s := m
	for i := len(a) - 1; i >= 0; i-- {
		m.Next = &ListNode{
			Val:  a[i],
			Next: nil,
		}
		m = m.Next
	}

	return s.Next

}

//leetcode submit region end(Prohibit modification and deletion)
