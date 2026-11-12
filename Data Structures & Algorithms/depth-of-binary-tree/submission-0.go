/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
   if root == nil {
	return 0
   }

   nodes := []*TreeNode{root} 
   maxDepth := 1

   for len(nodes) > 0 {
	if len(nodes) > maxDepth {
		maxDepth = len(nodes)
	}
	lastNode := nodes[len(nodes)-1]
	if lastNode.Left != nil {
		nodes = append(nodes, lastNode.Left)
		continue
	}
	if lastNode.Right != nil {
		nodes = append(nodes, lastNode.Right)
		continue
	}
	// go up
	for len(nodes) > 0  {
		last := nodes[len(nodes)-1]
		nodes = nodes[0:len(nodes)-1]

		if len(nodes) == 0 {
			break
		}

		
		newLast := nodes[len(nodes)-1]
		if newLast.Right != nil && newLast.Right != last {
			nodes = append(nodes, newLast.Right)
			break
		}
	}	
   }
   return maxDepth
}
