//https://www.hackerrank.com/challenges/strange-code/problem

func strangeCounter(t int64) int64 {
	rem := int64(3)

	for t > rem {
		t = t-rem
		rem *= 2
	}

	res := rem - t + 1
	return res
}
