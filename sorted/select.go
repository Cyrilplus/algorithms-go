package sorted


func SelectSorted(data Interface) {
	length := data.Len()
	for i, max := 0, length - 1; i < max; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(minIndex, i)
	}
}
