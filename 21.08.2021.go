//https://www.hackerrank.com/challenges/minimum-distances/problem

func minimumDistances(a []int32) int32 {
	var arr []int32
	var min int32
	for i:= range a {
		for j := i+1; j <len(a); j++ {
			if a[i] == a[j] {
				min := j - i
				arr = append(arr, int32(min))
				break
			}
		}
	}
	if len(arr) <= 0 {
		min = -1
	}else {
		min = arr[0]
		for i:=1; i<len(arr); i++ {
			if min > arr[i] {
			min = arr[i]
			}
		}
	}
	return min
}