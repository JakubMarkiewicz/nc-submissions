/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	visited := make(map[*TreeNode]int)
	
	res := 0

	dfs := []*TreeNode{root}

	for len(dfs) > 0 {
		lastElem := dfs[len(dfs)-1]	

		lVal, lVisited := visited[lastElem.Left]	
		rVal, rVisited := visited[lastElem.Right]	

		if lastElem.Left != nil && !lVisited {
			dfs = append(dfs, lastElem.Left)
			continue
		}

		if lastElem.Right != nil && !rVisited {
			dfs = append(dfs,  lastElem.Right)
			continue
		}


		localMax := intMax(lVal,rVal) + 1

		visited[lastElem] = localMax

		pathLength := lVal + rVal
		if pathLength > res {
			res = pathLength
		}

		dfs = dfs[0:len(dfs)-1]
	}

	return res
}

func intMax(a,b int) int {
	if a > b {
		return a
	}

	return b
}
