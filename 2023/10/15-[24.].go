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
// Related Topics 递归 链表 👍 2057 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//     Val int
//    Next *ListNode
// }

func swapPairs(head *ListNode) *ListNode {
	// a->b->c->d
	// b->a->d->c

	//-------------------------------------------
	//if head==nil||head.Next==nil{
	//	return head
	//}
	//newHead:=head.Next   //先拿到 b
	//head.Next=swapPairs(newHead.Next) //处理b后面的节点，并让 a 指向 后面交换后的节点 // a->d
	//newHead.Next=head //交换 a b // b 指向 a
	//return newHead
	//------------------------------------

	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil { //0->a->b->c->d
		node1 := temp.Next      //a
		node2 := temp.Next.Next //b
		//交换前面两个节点
		temp.Next = node2       // 0->b
		node1.Next = node2.Next // a->c
		node2.Next = node1      // b->a
		// 0->b->a->c->d
		temp = node1 // a 赋给 temp //已处理两个节点，将处理下两个节点
	}

	return dummyHead.Next

}

//leetcode submit region end(Prohibit modification and deletion)
