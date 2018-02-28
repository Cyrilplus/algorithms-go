package sorted

import (
	"sort"
	"testing"
)

type Arr []int

func (arr Arr) Len() int           { return len(arr) }
func (arr Arr) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }
func (arr Arr) Less(i, j int) bool { return arr[i] < arr[j] }

func TestSelectSorted(t *testing.T) {
	testArr := Arr([]int{4, 3, 2, 0, 10, 23, 5, 0, 1, -10})
	t.Log(testArr)
	SelectSorted(testArr)
	t.Log(testArr)
	if sort.IsSorted(testArr) {
		t.Log("successfully")
	} else {
		t.Log(testArr)
		t.Error("failed")
	}
}

func TestInsertSorted(t *testing.T) {
	testArr := Arr([]int{4, 3, 2, 0, 10, 23, 5, 0, 1, -10})
	t.Log(testArr)
	InsertSorted(testArr)
	t.Log(testArr)
	if sort.IsSorted(testArr) {
		t.Log("successfully")
	} else {
		t.Log(testArr)
		t.Error("failed")
	}
}

func TestShellSorted(t *testing.T) {
	testArr := Arr([]int{4, 3, 2, 0, 10, 23, 5, 0, 1, -10})
	t.Log(testArr)
	ShellSorted(testArr)
	t.Log(testArr)
	if sort.IsSorted(testArr) {
		t.Log("successfully")
	} else {
		t.Log(testArr)
		t.Error("failed")
	}
}

func TestMergeSorted(t *testing.T) {
	testArr := Arr([]int{4, 3, 2, 0, 10, 23, 5, 0, 1, -10})
	t.Log(testArr)
	MergeSorted(testArr, 0, testArr.Len() - 1)
	t.Log(testArr)
	if sort.IsSorted(testArr) {
		t.Log("successfully")
	} else {
		t.Log(testArr)
		t.Error("failed")
	}
}
