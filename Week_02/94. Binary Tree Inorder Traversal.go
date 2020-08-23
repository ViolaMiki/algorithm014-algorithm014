package Week_02

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	resRight := inorderTraversal(root.Right)

	resLeft := inorderTraversal(root.Left)

	return append(append(resLeft, root.Val), resRight...)
}
