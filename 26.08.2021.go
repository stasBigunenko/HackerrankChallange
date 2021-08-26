//https://www.hackerrank.com/challenges/runningtime/problem

func runningTime(arr []int32) int32 {
	var count int32 = 0

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j]{
				count++
			}
		}
	}
	return count
}
