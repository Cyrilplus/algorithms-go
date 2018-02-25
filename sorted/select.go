package sorted


func SelectSorted(data Interface) {
	length := data.Len() - 1
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if data.Less(j + 1, minIndex) {
				minIndex = j + 1
			}
		}
		data.Swap(minIndex, i)
	}
}

func InsertSorted(data Interface) {
	
}