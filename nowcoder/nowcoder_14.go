package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表中倒数第K个节点
// 链表 -> 双指针
func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	if pListHead == nil {
		return nil
	}

	cur := &ListNode{
		Next: pListHead,
	}
	pre := cur
	for i := 0; i < k && cur != nil; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return nil
	}

	for cur != nil {
		cur = cur.Next
		pre = pre.Next
	}

	return pre
}
