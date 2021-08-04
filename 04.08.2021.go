//https://www.hackerrank.com/challenges/equality-in-a-array/problem

func equalizeArray(arr []int32) int32 {
	var num, sum1, sum2 int32
	sum2 = 0

	for i:=0; i<len(arr); i++ {
		sum1 = 1
		for j:=i+1; j<len(arr); j++ {
			if arr[i] == arr[j] {
				sum1++
			}
		}
		if sum2 < sum1 {
			sum2 = sum1
		}
	}
	qty := int32(len(arr)) - sum2
	return qty
}
