/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
   curr := &ListNode{Val: 0}
   res := ListNode{Next: curr}

   for l1 != nil || l2 != nil {
	newNode := ListNode{Val: 0}
	if l1 != nil {
		newNode.Val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		newNode.Val += l2.Val
		l2 = l2.Next
	}

	if curr.Val >= 10  {
		newNode.Val += 1
		curr.Val = curr.Val % 10
	}

	curr.Next = &newNode
	curr = &newNode
   }

	if curr.Val >= 10  {
		newNode := &ListNode{Val: 1}
		curr.Val = curr.Val % 10
		curr.Next = newNode
	}

   return res.Next.Next
}
