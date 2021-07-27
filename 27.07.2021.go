//https://www.hackerrank.com/challenges/cats-and-a-mouse/problem

func catAndMouse(x int32, y int32, z int32) string {
	var m string
	var catA, catB int32

	catA = x - z
	catB = y - z

	if catA < 0 {
		catA *= -1
	}

	if catB < 0 {
		catB *= -1
	}

	if catA < catB {
		m = "Cat A"
	}else if catA > catB {
		m = "Cat B"
	}else {
		m = "Mouse C"
	}

	return m
}
