package main

/*
反转一个单链表。

进阶:
链表可以迭代或递归地反转。你能否两个都实现一遍？
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	currentNode := head
	var preNode *ListNode
	preNode = nil
	nextNode := currentNode.Next
	for currentNode != nil && nextNode != nil {
		currentNode.Next = preNode
		preNode = currentNode
		currentNode = nextNode
		nextNode = currentNode.Next
	}
	currentNode.Next = preNode
	return currentNode

}

type ListNode struct {
	Val  int
	Next *ListNode
}


// simple
func reverseListSimple(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
