package leetcode

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 2545 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//	type ListNode struct {
//	   Val int
//	   Next *ListNode
//	}
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	//m := head
	//p := &ListNode{
	//	Val:  0,
	//	Next: nil,
	//}
	//
	//pp := p
	//
	//var a, b int
	//for head != nil {
	//	a++
	//	head = head.Next
	//}
	//for m != nil {
	//
	//	b++
	//	if b != a-n+1 {
	//		p.Next = &ListNode{
	//			Val:  m.Val,
	//			Next: nil,
	//		}
	//		p = p.Next
	//	}
	//	m = m.Next
	//}
	//
	//return pp.Next

	var length int
	for head != nil {
		length++
		head = head.Next
	}

	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next

}

//leetcode submit region end(Prohibit modification and deletion)
