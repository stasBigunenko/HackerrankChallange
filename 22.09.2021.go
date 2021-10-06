//https://www.hackerrank.com/challenges/between-two-sets/problem

func getTotalX(a []int32, b []int32) int32 {
	flag1 := 0
	flag2 := 0
	res := int32(0)

	for k := int32(1); k <= 100; k++ {
		for i := 0; i < len(a); i++ {
			if k % a[i] == 0 {
				flag1 = 1
			} else {
				flag1 = 0
				break
			}
		}
		if flag1 == 1 {
			for j := 0; j < len(b); j++ {
				if b[j] % k == 0 {
					flag2 = 1
				} else {
					flag2 = 0
					break
				}
			}
			if flag2 == 1 {
				res++
			}
		}
	}
	return res
}