//https://www.hackerrank.com/challenges/encryption/problem

//needed additional import math to find square root

func encryption(s string) string {
	var row, col int
	var res string
	var s1 []string

	sqr := math.Sqrt(float64(len(s)))
	intpart, decpart := math.Modf(sqr)
	if decpart != 0 {
		row = int(sqr)
		col = int(sqr+1)
		if col * row < len(s) {
            		row++
        	}
	}else {
		row = int(intpart)
		col = int(intpart)
	}

	s1 = strings.Split(s, "")

	for i:=0; i<col; i++ {
		k := i
		for j:=0; j<row; j++ {
			if k < len(s) {
				res += s1[k]
				k += col
			}
		}
		res += " "
	}
	return res
}
