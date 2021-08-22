//https://www.hackerrank.com/challenges/sherlock-and-squares/problem

func squares(a int32, b int32) int32 {
	var count int32 = 0
	for i:=int32(1); i<=b/2; i++ {
		if i * i > b {
			break
		}
		if i * i >= a && i * i <= b {
			count++
		}
	}
	return count
}