package qs

import "testing"

func TestPartitionOne(t *testing.T) {
	testData := []int{26, 36, 89, 80, 87, 66, 86, 49, 56, 14}
	p := partition(testData, 0, len(testData)-1)
	if p != 0 {
		t.Error("Returned partition location isn't correct")
	}
}

func TestPartitionTwo(t *testing.T) {
	testData := []int{14, 26, 36, 89, 80, 87, 66, 86, 49, 56}
	p := partition(testData, 0, len(testData)-1)
	if p != 4 {
		t.Error("Returned partition location isn't correct")
	}
}

func TestPartitionThree(t *testing.T) {
	testData := []int{14, 26, 36, 89, 80, 87, 66, 86, 56, 49}
	p := partition(testData, 0, len(testData)-1)
	if p != 3 {
		t.Error("Returned partition location isn't correct")
	}
}

func TestSortSimple(t *testing.T) {
	testData := []int{3, 2, 1}
	QuickSort(testData, 0, len(testData)-1)

	for i := 1; i < len(testData); i++ {
		if testData[i-1] > testData[i] {
			t.Error("Data isn't correctly ordered:", testData[i-1], ">", testData[i])
		}
	}

}

func TestSortLong(t *testing.T) {
	testData := []int{26, 36, 89, 80, 87, 66, 86, 49, 56, 14}
	QuickSort(testData, 0, len(testData)-1)

	for i := 1; i < len(testData); i++ {
		if testData[i-1] > testData[i] {
			t.Error("Data isn't correctly ordered:", testData[i-1], ">", testData[i])
		}
	}

}
