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
	visited := make(map[*TreeNode]struct{})

	max := 0
	dfs := []*TreeNode{root}

	for len(dfs) > 0 {
		lastElem := dfs[len(dfs)-1]
		visited[lastElem] = struct{}{}

		if len(dfs) > max {
			max = len(dfs)
		}

		_, lVisited := visited[lastElem.Left]
		_, rVisited := visited[lastElem.Right]

		if lastElem.Left != nil && !lVisited  {
			dfs = append(dfs, lastElem.Left)
			continue
		}

		if lastElem.Right != nil && !rVisited {
			dfs = append(dfs, lastElem.Right)
			continue
		}

		dfs = dfs[0:len(dfs)-1]
	}

	return max
}
