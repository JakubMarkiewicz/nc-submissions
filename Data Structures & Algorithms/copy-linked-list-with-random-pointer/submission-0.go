/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	cp := make(map[*Node]*Node)

	tmp := head

	for tmp != nil {
		n := &Node{Val: tmp.Val}
		cp[tmp] = n	
		tmp = tmp.Next
	}

	tmp = head

	for tmp != nil {
		cpNode := cp[tmp]	
		cpNode.Next = cp[tmp.Next]
		cpNode.Random = cp[tmp.Random]
		tmp = tmp.Next
	}

	return cp[head]
}

