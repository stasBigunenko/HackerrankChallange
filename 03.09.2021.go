//https://www.hackerrank.com/challenges/sherlock-and-array/problem

func balancedSums(arr []int32) string {
	if len(arr) == 1 {
		return "YES"
	}

	var sum2 int32

	for _, val := range arr {
		sum2 += val
	}

	sum1 := int32(0)
	for i := 0; i < len(arr); i++ {
		sum2 -= arr[i]
		if sum1 < sum2 {
			sum1 += arr[i]
		}else if sum1 == sum2 {
			return "YES"
		}
	}
	return "NO"
}