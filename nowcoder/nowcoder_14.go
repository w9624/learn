package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	// write code here
	if pListHead == nil {
		return pListHead
	}

	head := &ListNode{
		Next: pListHead,
	}
	cur := head
	for i := 0; i < k && cur != nil; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return nil
	}

	pre := head
	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}

	return pre
}
