//https://www.hackerrank.com/challenges/find-digits/problem

func findDigits(n int32) int32 {
	var digit int32
	var count int32 = 0
	convert := strconv.Itoa(int(n))
	for _, val:= range convert {
		digit = int32(val - '0')
		if digit != 0 {
			if n % digit == 0 {
				count++
			}
		}
	}
	return count
}
