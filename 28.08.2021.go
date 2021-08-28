//https://www.hackerrank.com/challenges/funny-string/problem

func funnyString(s string) string {
	var srev string
	var dif int
	var tmp int
	str := "Funny"

	last := len(s)-1
	for i:= last; i >=0; i-- {
		srev += string(s[i])
	}

	for i:=1; i < len(s); i++ {
		tmp = int(s[i-1]) - int(s[i])
		if tmp < 0 {
			tmp *= -1
		}
		dif = int(srev[i-1]) - int(srev[i])
		if dif < 0 {
			dif *= -1
		}

		if dif != tmp {
			str = "Not Funny"
			break
		}
	}
	return str
}