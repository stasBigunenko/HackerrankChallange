//https://www.hackerrank.com/challenges/reduced-string/problem

func superReducedString(s string) string {
	splitStr := strings.Split(s, "")
	var res string

	for i:=0; i<len(splitStr)-1; i++{
		if splitStr[i] == splitStr[i+1] {
			splitStr = append(splitStr[:i], splitStr[i+2:]...)
			i = -1
		}
	}
	res = strings.Join(splitStr, "")
	if res == "" {
		res = "Empty String"
	}
	return res
}