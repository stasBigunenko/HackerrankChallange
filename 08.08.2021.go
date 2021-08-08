//https://www.hackerrank.com/challenges/taum-and-bday/problem

func taumBdawc(b int32, w int32, bc int32, wc int32, z int32) int64 {
	var cost int64 = 0
	var min1 int64
	var min2 int64
	if bc > wc + z {
		min1 = int64(wc+z)
	}else {
		min1 = int64(bc)
	}
	if wc > bc + z {
		min2 = int64(bc+z)
	}else {
		min2 = int64(wc)
	}
	cost = min1*int64(b) + min2*int64(w)
	return int64(cost)
}