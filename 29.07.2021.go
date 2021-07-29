//https://www.hackerrank.com/challenges/utopian-tree/problem

func utopianTree(n int32) int32 {
	height := int32(1)

	for i:=int32(1); i<=n; i++ {
		if i % 2 != 0 {
			height *= int32(2)
		}else{
			height += int32(1)
		}
	}
	return height
}