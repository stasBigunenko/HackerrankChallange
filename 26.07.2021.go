//Electronics Shop

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	max := int32(-1)

	for i:=0; i < len(keyboards); i++ {
		for j:=0; j< len(drives); j++ {
			if (keyboards[i] + drives[j] <= b) && (max <= keyboards[i] + drives [j]) {
				max = keyboards[i] + drives[j]
			}
		}
	}
	return max
}