/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	visited := map[*ListNode]struct{}{}

	for head != nil {
		if _, ok := visited[head]; ok {
			return true	
		}
		visited[head] = struct{}{}
		head = head.Next
	}

	return false
}
