//https://www.hackerrank.com/challenges/append-and-delete/problem

func appendAndDelete(s string, t string, k int32) string {
	count := 0
	res := "No"

	if len(s) <= len(t) {
		for i := range s {
			if s[i] == t[i] {
				count++
			}else{
				break
			}
		}
	}else {
		for i := range t {
			if s[i] == t[i] {
				count++
			}else{
				break
			}
		}
	}

	sum1 := len(s)-count
	sum2 := len(t)-count

	if (sum1 + sum2 <= int(k) && (sum1 + sum2) % 2 == int(k % 2)) || len(s) + len(t) < int(k) {
		res = "Yes"
	}
	return res
}