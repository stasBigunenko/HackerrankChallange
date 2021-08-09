//https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem

func jumpingOnClouds(c []int32, k int32) int32 {
	var energy int32 = 100
	for i:=int32(0); i<int32(len(c)); i += k {
		if c[i] == 1 {
			energy -= 3
		}else {
			energy -= 1
		}
		if i + k >= int32(len(c)){
			if (i + k - int32(len(c)) == 0 ){
				break
			}else{
				i -= int32(len(c))
			}
		}
	}
	return energy
}
