/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	next := res


	for list1 != nil || list2 != nil {
		if list2 == nil || (list1 != nil && list1.Val < list2.Val)  {
			tmp := &ListNode{Val: list1.Val}
			next.Next = tmp
			list1 = list1.Next
		} else {
			tmp := &ListNode{Val: list2.Val}
			next.Next = tmp
			list2 = list2.Next
		}
		next = next.Next
	}

	return res.Next
}
