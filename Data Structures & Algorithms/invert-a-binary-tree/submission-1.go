/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
if root == nil {
	return nil
}

  bfs := []*TreeNode{root}  

  for len(bfs) > 0 {
	firstElem := bfs[0]
	firstElem.Left, firstElem.Right = firstElem.Right, firstElem.Left

	if firstElem.Left != nil {
		bfs = append(bfs,firstElem.Left)
	}
	if firstElem.Right != nil {
		bfs = append(bfs,firstElem.Right)
	}
	bfs = bfs[1:]
  }

  return root
}
