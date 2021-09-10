//https://www.hackerrank.com/challenges/manasa-and-stones/problem

func stones(n int32, a int32, b int32) []int32 {
	var res []int32

	if a == b {
		res = append(res, a * (n-1))
	}else {
		if a > b {
			a, b = b, a
		}
		for i:=n-1; i>=0; i-- {
			num := a * i + (n-i-1) * b
			res = append(res,num)
		}
	}
	return res
}
