package Week_02

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	resRight := preorderTraversal(root.Right)

	resLeft := preorderTraversal(root.Left)

	return append(append(append(res, root.Val), resLeft...), resRight...)
}
