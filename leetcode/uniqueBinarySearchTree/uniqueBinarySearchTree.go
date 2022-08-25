package uniquebinarysearchtree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTree(root int, allLeftTree, allRightTree []*TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0, len(allLeftTree)*len(allRightTree))
	for _, leftNode := range allLeftTree {
		for _, rightNode := range allRightTree {
			rootNode := &TreeNode{
				Val:   root,
				Left:  leftNode,
				Right: rightNode,
			}
			res = append(res, rootNode)
		}
	}
	return res
}

func buildTree(l, r int) []*TreeNode {
	if l == r {
		if l == -1 || r == -1 {
			return []*TreeNode{nil}
		} else {
			return []*TreeNode{
				&TreeNode{
					Val: l,
				},
			}
		}
	}

	allTree := []*TreeNode{}
	for root := l; root <= r; root++ {
		leftTree := []*TreeNode{}
		rightTree := []*TreeNode{}
		if root == l {
			leftTree = buildTree(-1, -1)
		} else {
			leftTree = buildTree(l, root-1)
		}

		if root == r {
			rightTree = buildTree(-1, -1)
		} else {
			rightTree = buildTree(root+1, r)
		}

		treeAtRoot := mergeTree(root, leftTree, rightTree)
		allTree = append(allTree, treeAtRoot...)
	}
	return allTree
}
func generateTrees(n int) []*TreeNode {
	return buildTree(1, n)
}
