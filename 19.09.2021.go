//https://www.hackerrank.com/challenges/picking-numbers/problem

func pickingNumbers(a []int32) int32 {
	var val int32
	var max int32 = 1

	for i :=0; i < len(a); i++ {
		var count int32 = 1
		val = a[i] - int32(1)

		for j := 0; j<len(a); j++ {
			if i != j {
				if a[j] == a[i] || a[j] == val {
					count++
				}
			}
		}
		if max < count {
			max = count
		}
	}
	return max
}
