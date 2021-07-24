//https://www.hackerrank.com/challenges/migratory-birds/problem

func migratoryBirds(arr []int32) int32 {
	var counter = []int32{0, 0, 0, 0, 0, 0,}
	result := int32(1)

	for i := range arr {
		counter[arr[i]]++;
	}

	max := counter[1]

	for i := int32(2); i < int32(6); i++ {
		if counter[i] > max {
			result = i;
			max = counter[i];
		}
	}

	return result;
}