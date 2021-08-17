//https://www.hackerrank.com/challenges/fair-rations/problem

func fairRations(B []int32) string {
	var count int32 = 0
	var m string
	for i, val := range B {
		if val % 2 != 0 && i+1 < len(B) {
			B[i] += 1
			B[i+1] += 1
			count+=2
		}
	}
	if B[len(B)-1] % 2 == 0 {
		m = strconv.Itoa(int(count))
	}else {
		m = "NO"
	}
	return m
}
