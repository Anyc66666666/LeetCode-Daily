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
// Related Topics 递归 链表 👍 3383 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//  Val int
//  Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	//list:=&ListNode{
	//	Val:  0,
	//	Next: nil,
	//}
	//m:=list
	//var a []int
	//
	//for head!=nil{
	//	a=append(a,head.Val)
	//	head=head.Next
	//}
	//
	//for i:=len(a)-1;i>=0;i--{
	//	list.Next=&ListNode{
	//		Val:  a[i],
	//	}
	//	list=list.Next
	//}
	//
	//return m.Next

	//----------------------------------------------

	//var prev *ListNode
	//curr:=head
	//for curr!=nil{  //  a b c
	//	next:=curr.Next // 先提前拿到 b
	//	curr.Next=prev // 由于反转，所以让 a 指向 a的前面
	//	prev=curr //由于for循环要往下走，a即将成为下一个b的prev
	//	curr=next //同理，curr要往下走到 b
	//}
	//return prev
	// O(n) O(1)

	//---------------------------------------------

	if head == nil || head.Next == nil { //递归最开始直接拿到倒数第二个节点
		return head
	}
	//a->b->c->d->e
	newHead := reverseList(head.Next) // a->b->c->d->e //now at d
	head.Next.Next = head             // 让 e->d  // d.Next 为 e , e.Next指向d
	head.Next = nil                   //再让 d指向 空 ，此时相当于在原来到基础上 翻转了 d e
	return newHead                    // a->b->c->d<-e

}

//leetcode submit region end(Prohibit modification and deletion)
