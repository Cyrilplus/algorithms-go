package sorted


func merge(data Interface, lo, mid, hi int) {
	// can not copy data, so instead of using insertion sort
	for i := mid + 1; i <= hi; i++ {
		for j := i - 1; j >= lo && data.Less(j + 1, j); j-- {
			data.Swap(j + 1, j)
		}
	}
}

func MergeSorted(data Interface, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	MergeSorted(data, lo, mid)
	MergeSorted(data, mid + 1, hi)
	merge(data, lo, mid, hi)
}

func MergeSortedNRecursion(data Interface) {
	N := data.Len()
	for l := 1; l < N; l *= 2 {
		for i, max := 0, N - l; i < max; i += l * 2 {
			hi := i + 2 * l - 1
			if N - 1 < hi {
				hi = N - 1
			}
			merge(data, i, i + l - 1, hi)
		}
	}
}