/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
   nodes := []*ListNode{}
   tmp := head 

   for tmp != nil {
	nodes = append(nodes, tmp)	
	tmp = tmp.Next
   }

	toRemoveIdx := len(nodes) - n

	if toRemoveIdx == 0 {
		head = head.Next
		return head
	}

	nodes[toRemoveIdx-1].Next = nodes[toRemoveIdx].Next

	return head
}
