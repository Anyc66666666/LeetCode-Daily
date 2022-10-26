package leetcode

import "fmt"

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 2756 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var l []int
	l1, l2 := list1, list2
	for {
		if l1 != nil {
			l = append(l, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			l = append(l, l2.Val)
			l2 = l2.Next

		}

		if l1 == nil && l2 == nil {
			break
		}
	}
	if len(l) == 0 {
		return nil
	}

	for i := 0; i < len(l)-1; i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i] > l[j] {
				b := l[i]
				l[i] = l[j]
				l[j] = b

			}
		}
	}
	fmt.Println(l)
	l3 := &ListNode{Val: l[0]}
	t := l3
	for i := 1; i < len(l); i++ {
		t.Next = &ListNode{Val: l[i]}
		t = t.Next

	}
	return l3

}

//leetcode submit region end(Prohibit modification and deletion)
