//https://www.hackerrank.com/challenges/beautiful-triplets/problem

func beautifulTriplets(d int32, arr []int32) int32 {
	var count int32 = 0

	for i := 0; i<len(arr); i++ {
		for j:=i; j<len(arr); j++ {
			if arr[j] - arr[i] == d {
				for k:=j; k<len(arr); k++ {
					if arr[k] - arr[j] == d {
						count++
					}
				}
			}
		}
	}
	return count
}