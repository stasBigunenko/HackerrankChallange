//https://www.hackerrank.com/challenges/day-of-the-programmer/problem

func dayOfProgrammer(year int32) string {
	y := strconv.Itoa(int(year))
	var m string
	if year == 1918 {
		m = "26.09.1918"
	}else if (year % 4 == 0 && (year < 1918 || year % 100 != 0)) || year % 400 == 0  {
		m = "12.09." + y
	}else {
		m = "13.09." + y
	}
	return m
}
