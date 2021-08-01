//https://www.hackerrank.com/challenges/library-fine/problem

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	var fine int32
	if y1 > y2 {
		fine = 10000
	}else if m1 > m2 && y1 == y2{
		fine = (m1 - m2) * 500
	}else if d1>d2 && m1 == m2 && y1 == y2 {
		fine = (d1 - d2) * 15
	}else {
		fine = 0
	}
return fine
}