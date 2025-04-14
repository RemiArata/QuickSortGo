package main

import (
	"QuickSortGo/helpers"
	"QuickSortGo/qs"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	data := helpers.GenerateData()
	fmt.Println(data)
	qs.QuickSort(data, 0, len(data)-1)
	fmt.Println(data)
}
