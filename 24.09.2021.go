//https://www.hackerrank.com/challenges/kaprekar-numbers/problem

//additional import "math"
func kaprekarNumbers(p int32, q int32) {
	count := 0
	res := ""
	var num1String, num2String []string
	var num1, num2 int
	for i:=p; i<=q; i++ {
		d := len(strconv.Itoa(int(i)))
		square := math.Pow(float64(i), 2)
		squareString := strconv.Itoa(int(square))
		l := len(squareString)
		splitSquare := strings.Split(squareString, "")
		if l <= 1 {
			if i == 1 || i == 9 {
				res += strconv.Itoa(int(i)) + " "
				count++
			}
		}else {
			num2String = splitSquare[l-d:]
			num2justString := strings.Join(num2String, "")
			num1String = splitSquare[:l-d]
			num1justString := strings.Join(num1String, "")
			num1, _ = strconv.Atoi(num1justString)
			num2, _ = strconv.Atoi(num2justString)
			if num1+num2 == int(i) {
				res += strconv.Itoa(int(i)) + " "
				count++
			}
		}
	}
	if count == 0 {
		res ="INVALID RANGE"
	}
	fmt.Println(res)
	return
}