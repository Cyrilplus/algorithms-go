package sorted

func ShellSorted(data Interface) {
	N := data.Len()
	h := 1
	for ; h < N / 3; {
		h = 3 * h + 1
	}
	for ; h >= 1; {
		for i := h; i < N; i++ {
			for j := i - h; j >=0 && data.Less(j + h, j); j -= h {
				data.Swap(j + h, j)
			}
		}
		h /= 3
	}
}

