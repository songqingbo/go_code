package main

/*

将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */

// 一波过
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		return &ListNode{l1.Val, mergeTwoLists(l1.Next, l2)}
	}
	return &ListNode{l2.Val, mergeTwoLists(l1, l2.Next)}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 一波过
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	move := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			move.Next = l1
			l1 = l1.Next
		} else {
			move.Next = l2
			l2 = l2.Next
		}
		move = move.Next
	}
	if l1 == nil {
		move.Next = l2
		return head.Next
	}
	move.Next = l1
	return head.Next
}
