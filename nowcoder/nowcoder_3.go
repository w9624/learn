package main

import (
	"log"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 从尾到头打印链表
// 链表 -> 头插法/逆置
func PrintListFromTailToHead(listNode *ListNode) (array []int) {

	if listNode == nil {
		return array
	}

	head := &ListNode{}

	for cur := listNode; cur != nil; {
		// new node
		node := &ListNode{
			Val: cur.Val,
		}

		// Next
		cur = cur.Next

		// insert
		node.Next = head.Next
		head.Next = node
	}

	head = head.Next
	for cur := head; cur != nil; cur = cur.Next {
		array = append(array, cur.Val)
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
				cur := head.Next
				for _, i := range []int{67, 0, 24, 58} {
					cur = &ListNode{
						Val:  i,
						Next: nil,
					}
					pre.Next = cur

					cur = cur.Next
					pre = pre.Next
				}

				return head.Next
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
