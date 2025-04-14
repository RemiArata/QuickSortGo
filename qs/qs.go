package qs

func QuickSort(data []int, low int, high int) {
	if low >= high || low < 0 {
		return
	}

	p := partition(data, low, high)

	QuickSort(data, low, p-1)  // bottom
	QuickSort(data, p+1, high) // top
}

func partition(data []int, low int, high int) int {
	pivot := data[high]
	i := low

	for j := low; j < high; j++ {
		if data[j] < pivot {
			data[j], data[i] = data[i], data[j]
			i = i + 1
		}
	}
	data[i], data[high] = data[high], data[i]
	return i
}

//                        h     p
// [28 34 29 21 23 40 38 84 78 61]
