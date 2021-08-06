//https://www.hackerrank.com/challenges/lisa-workbook/problem

func workbook(n int32, k int32, arr []int32) int32 {
	var p int32
	var res int32 = 0
	for _, val := range arr {
		p += 1
		for j:=int32(1); j<=val; j++ {
			if j == p {
				res++
			}
			if j == val {
				continue
			}
			if j % k == 0 {
				p += 1
			}
		}
	}
	return res
}