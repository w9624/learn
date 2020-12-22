/*
 * Copyright (C) Runze Wang
 * Copyright (C) ixiaochuan.cn
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// step 1: 将两个链表表反转，双指针同时从头部出发，直到两个链表的节点地址不同（需改动链表）T(n) S(1)
// step 2: 使用栈，将两个链表入栈，然后同时出栈比较，和step1逻辑类似 （无需改动链表）T(n) S(m+n)
// step 3: 先计算两个链表长度p1/p2，然后长的链表向后走abs(p1-p2)位置，接着同时向后走直到出现相同节点 T(n)
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	var p1Len, p2Len int
	p1 := pHead1
	for p1 != nil {
		p1 = p1.Next
		p1Len++
	}
	p2 := pHead2
	for p2 != nil {
		p2 = p2.Next
		p2Len++
	}

	for p1Len > p2Len {
		p1Len--
		pHead1 = pHead1.Next
	}
	for p2Len > p1Len {
		p2Len--
		pHead2 = pHead2.Next
	}

	var node *ListNode
	for pHead1 != nil && pHead2 != nil {
		if pHead1 == pHead2 {
			node = pHead1
			break
		}
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next
	}

	return node
}

func main() {

}
