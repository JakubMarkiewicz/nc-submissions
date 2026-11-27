/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	visited := make(map[int]struct{})

	stackP, stackQ := []*TreeNode{p}, []*TreeNode{q}

	for len(stackP) > 0 || len(stackQ) > 0{
		pNode, qNode := stackP[len(stackP)-1], stackQ[len(stackQ)-1]
		if pNode == nil || qNode == nil || pNode.Val != qNode.Val {
			return false
		}

		if _, ok := visited[pNode.Val]; !ok {
			visited[pNode.Val] = struct{}{}
			if  pNode.Left != nil || qNode.Left != nil {
				stackP, stackQ = append(stackP, pNode.Left), append(stackQ, qNode.Left)
			}

			if pNode.Right != nil || qNode.Right != nil {
				stackP, stackQ = append(stackP, pNode.Right), append(stackQ, qNode.Right)
			}
			continue
		}

		stackP, stackQ = stackP[0:len(stackP)-1], stackQ[0:len(stackQ)-1]
	} 


   return len(stackP) == len(stackQ)
}
