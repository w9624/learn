package main

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

// 复制复杂链表
// 链表 -> 复制 -> 重点是先复制后拆分
func Clone(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		node := &RandomListNode{
			Label:  cur.Label,
			Next:   cur.Next,
			Random: cur.Random,
		}
		cur.Next = node
		cur = node.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	headNew := &RandomListNode{}
	pre := headNew
	for cur != nil {
		pre.Next = cur.Next
		pre = pre.Next
		cur = cur.Next.Next
	}

	return headNew.Next
}
