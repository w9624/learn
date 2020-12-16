package main

import (
	"log"
	"reflect"
)

type ListNode struct {
	val  int
	next *ListNode
}

// 从尾到头打印链表
func PrintListFromTailToHead(listNode *ListNode) (array []int) {

	if listNode == nil {
		return array
	}

	head := &ListNode{}

	for cur := listNode; cur != nil; {
		// new node
		node := &ListNode{
			val: cur.val,
		}

		// next
		cur = cur.next

		// insert
		node.next = head.next
		head.next = node
	}

	head = head.next
	for cur := head; cur != nil; cur = cur.next {
		array = append(array, cur.val)
	}

	return array
}

func main() {
	cases := []struct {
		listNode func() *ListNode
		expected []int
	}{
		{
			listNode: func() *ListNode {

				head := &ListNode{}
				pre := head
				cur := head.next
				for _, i := range []int{67, 0, 24, 58} {
					cur = &ListNode{
						val:  i,
						next: nil,
					}
					pre.next = cur

					cur = cur.next
					pre = pre.next
				}

				return head.next
			},
			expected: []int{58, 24, 0, 67},
		},
	}

	for _, c := range cases {
		result := func(listNodeFunc func() *ListNode, expected []int) bool {

			listNode := listNodeFunc()

			log.Printf("do print")
			array := PrintListFromTailToHead(listNode)
			log.Printf("array: %v", array)

			return reflect.DeepEqual(array, expected)
		}(c.listNode, c.expected)

		if !result {
			log.Fatalf("fail")
		}
	}

}
