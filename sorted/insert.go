package sorted

func InsertSorted(data Interface) {
	length := data.Len()
	for i := 1; i < length; i++ {
		for j := i - 1; j >=0; j-- {
			if data.Less(j + 1, j) {
				data.Swap(j + 1, j)
			}
		} 
	}
}
