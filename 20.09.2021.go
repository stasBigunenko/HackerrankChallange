//https://www.hackerrank.com/challenges/closest-numbers/problem

//addition import "sort"

func closestNumbers(arr []int32) []int32 {
	var min int32
	var res []int32
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	min = arr[1] - arr[0]
	for i := 2; i<len(arr); i++ {
		if min > arr[i] - arr[i-1] {
			min = arr[i] - arr[i-1]
		}
	}

	for i := 1; i<len(arr); i++ {
		if min == arr[i] - arr[i-1] {
			res = append(res, arr[i-1])
			res = append(res, arr[i])
		}
	}
	return res
}
