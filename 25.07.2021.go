//https://www.hackerrank.com/challenges/counting-valleys/problem

func countingValleys(steps int32, path string) int32 {
	valley := int32(0)
	count := int32(0)

	for i := range path {
		if path[i] == 85 {
			count++
		}else{
			count--
		}

		if count ==0 && path[i] == 85 {
			valley++
		}
	}
	return valley
}