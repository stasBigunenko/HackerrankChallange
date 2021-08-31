//https://www.hackerrank.com/challenges/lonely-integer/problem

func lonelyinteger(a []int32) int32 {
	b := [100]int32{}
	var res int32

	for _, val := range a {
		b[val]++
	}

	for i := range b {
		if b[i] == int32(1) {
			res = int32(i)
			break
		}
	}
	return res
}
