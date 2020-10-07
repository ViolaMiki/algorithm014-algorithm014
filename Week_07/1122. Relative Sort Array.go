package Week_07

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := []int{}
	arr2Map := map[int][]int{}
	arr4 := []int{}

	for _, i := range arr2 {
		arr2Map[i] = make([]int, 0)
	}

	for _, i := range arr1 {
		if _, ok := arr2Map[i]; ok {
			arr2Map[i] = append(arr2Map[i], i)
		} else {
			arr4 = append(arr4, i)
		}
	}

	for _, i := range arr2 {
		res = append(res, arr2Map[i]...)
	}

	sort.Ints(arr4)
	return append(res, arr4...)
}
