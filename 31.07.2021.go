//https://www.hackerrank.com/challenges/angry-professor/problem?h_r=next-challenge&h_v=zen

func angryProfessor(k int32, a []int32) string {
	var msg string = "YES"
	var count int32 = 0

	for _, val := range a {
		if val <= 0 {
			count++
		}
	}
	if count >= k {
		msg = "NO"
	}
	return msg
}