package Week_06

func generateParenthesis(n int) []string {
	res := []string{}

	appendParenthesis(n, n, "", &res)

	return res
}

func appendParenthesis(left int, right int, str string, res *[]string) {
	if right == 0 {
		*res = append(*res, str)
		return
	}

	if left > 0 {
		appendParenthesis(left - 1, right, str + "(", res)
	}

	if right > left {
		appendParenthesis(left, right - 1, str + ")", res)
	}
}
