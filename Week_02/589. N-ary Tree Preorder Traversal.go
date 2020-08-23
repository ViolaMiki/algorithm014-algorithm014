package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*Node{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}
