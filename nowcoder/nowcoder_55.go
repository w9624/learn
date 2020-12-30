package main

// 链表中环的入口节点
// 链表 -> 双指针
// 定理：没有这个没法用双指针解题。。。
// 1. fast走两步，snow走一步，如果存在环则fast和snow一定相遇；
// 2. 相遇后fast从头开始走，snow继续走，每次都是走一步，相遇在环开始点
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	fast, snow := pHead, pHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		snow = snow.Next

		if fast == snow {
			fast = pHead
			for fast != snow {
				fast = fast.Next
				snow = snow.Next
			}

			if fast == snow {
				return fast
			}
		}
	}

	return nil
}

func main() {

}
