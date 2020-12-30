package main

// 链表 -> 双指针法
// 记录头节点，从第一个节点向后比较，删除相同节点
func deleteDuplication(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	head := &ListNode{}
	pre := head
	pre.Next = pHead
	cur := pHead.Next
	for cur != nil {
		if pre.Next.Val == cur.Val {
			for cur != nil && pre.Next.Val == cur.Val {
				cur = cur.Next
			}
			pre.Next = cur
			if cur == nil {
				break
			}
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}

	return head
}
