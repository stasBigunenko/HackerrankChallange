//https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem

func minimumAbsoluteDifference(arr []int32) int32 {
	var min int32
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	min = arr[len(arr)-1]-arr[0]

	for i:=0; i<len(arr)-1; i++ {
		if min > (arr[i+1]-arr[i]) {
			min = arr[i+1]-arr[i]
		}
	}
	return min
}
