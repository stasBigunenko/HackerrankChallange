//https://www.hackerrank.com/challenges/designer-pdf-viewer/problem

func designerPdfViewer(h []int32, word string) int32 {
		var runes []int32
		var max, count int32

		runes = append(runes, 'a')

		for i:=int32(1); i<int32(len(h)); i++ {
			runes = append(runes, runes[0] + i)
		}

		for i,val := range h {
			for j := range word {
				if byte(val) == word[j] {
					count++
					if max <= h[i] {
						max = h[i]
					}
				}
			}
		}
	return max * count
}
