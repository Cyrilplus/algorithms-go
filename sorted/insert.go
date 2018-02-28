package sorted

func InsertSorted(data Interface) {
	length := data.Len()
	for i := 1; i < length; i++ {
		for j := i - 1; j >=0 && data.Less(j + 1, j); j-- {
				data.Swap(j + 1, j)
		} 
	}
}
