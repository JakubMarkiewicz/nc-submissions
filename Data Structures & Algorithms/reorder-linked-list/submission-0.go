/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
   if head == nil {
	return
   } 

   nodes := []*ListNode{}
   tmp := head

   for tmp != nil {
	nodes = append(nodes, tmp)
	tmp = tmp.Next
   }

   l, r := 0, len(nodes) - 1

   for l < r {
	nodes[l].Next = nodes[r]
	l++
	nodes[r].Next = nodes[l]
	r--
   }

   nodes[l].Next = nil

}
