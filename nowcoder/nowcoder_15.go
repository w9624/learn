package main

// 反转链表
// 链表 -> 头插法/逆置
func ReverseList(pHead *ListNode) *ListNode {

	head := &ListNode{}
	cur := pHead
	for cur != nil {
		node := head.Next
		head.Next = cur
		cur = cur.Next
		head.Next.Next = node
	}

	return head.Next
}
