//https://www.hackerrank.com/challenges/sock-merchant/problem

func sockMerchant(n int32, ar []int32) int32 {
	var pairs int32 = 0
	for i := range ar {
		sock := int32(1)

		if ar[i] != 0 {
		for j := i + 1; j < len(ar); j++ {
			if ar[i] == ar[j] {
				sock++
				ar[j] = 0
			}
		}
		pairs += sock / 2
		}
	}
	return pairs
}