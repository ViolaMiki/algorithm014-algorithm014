package Week_03

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	skip := 0
	for k, v := range inorder {
		if v == preorder[0] {
			skip = k
			break
		}
	}

	lTree := buildTree(preorder[1:skip + 1], inorder[0:skip])
	rTree := buildTree(preorder[skip + 1:], inorder[skip + 1:])

	root := &TreeNode{
		Val: preorder[0],
		Left: lTree,
		Right: rTree,
	}
	return root
}