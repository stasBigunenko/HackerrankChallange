//https://www.hackerrank.com/challenges/game-of-thrones/problem

func gameOfThrones(s string) string {
	var res string = "YES"
	count := [26]int{0}
	var flag int = 0

	for i := range s {
		count[s[i] - 'a']++
	}

	if len(s) % 2 != 0 {
		for _, val := range count {
			if val % 2 != 0 {
				flag++
			}
			if flag > 1 {
				res = "NO"
				break
			}
		}
	}else {
		for _, val := range count {
			if val % 2 != 0 {
				res = "NO"
				break
			}
		}
	}
	return res
}