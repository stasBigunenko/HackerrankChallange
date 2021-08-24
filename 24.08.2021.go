//https://www.hackerrank.com/challenges/strong-password/problem

func minimumNumber(n int32, password string) int32 {
	flag := 4
	for _, val := range password {
		if val >= '0' && val <= '9' {
			flag -=1
			break
		}
	}
	for _, val := range password {
		//fmt.Println(val)
		if val >= 'A' && val <= 'Z' {
			flag -= 1
			break
		}
	}
	for _, val := range password {
		if val >= 'a' && val <= 'z' {
			flag -=1
			break
		}
	}
	for _, val := range password {
		if (val >= '#' && val <= '&') || (val >= '(' && val <= '+') || val == '@' || val == '^' || val == '!' || val == '-' {
			flag -=1
			break
		}
	}
	if len(password) < 6 {
		if 6 - len(password) > flag {
			flag = 6 - len(password)
		}
	}
	res := int32(flag)
	return res
}
