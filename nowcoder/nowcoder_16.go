package main

// 合并两个排序链表
// 链表 -> 合并
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {

	head := &ListNode{}
	cur := head
	p1, p2 := pHead1, pHead2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}

	if p1 != nil {
		cur.Next = p1
	}

	if p2 != nil {
		cur.Next = p2
	}

	return head.Next
}
