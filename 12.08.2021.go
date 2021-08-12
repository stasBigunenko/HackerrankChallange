//https://www.hackerrank.com/challenges/encryption/problem

func encryption(s string) string {
	var row, col int
	var res string
	var s1 []string
	
	i := len(s)/2
	for ; i>0; i-- {
		if i * i <= len(s) {
			break
		}
	}
	
	if i * i == len(s) {
		row = i
		col = i
	}else {
		row = i
		col = i + 1
		if row * col < len(s) {
			row++
		}
	}

	s1 = strings.Split(s, "")
	
	for i=0; i<col; i++ {
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
