package sorted

func merge(data Interface, lo, mid, hi int) {
	// 因为无法复制数组，这个地方采用了插入排序作为归并排序
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