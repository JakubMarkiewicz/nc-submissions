/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	var val *ListNode
	for head != nil {
		node := ListNode{Val: head.Val, Next: val}
		val = &node
		head = head.Next
	}

	return val
}
