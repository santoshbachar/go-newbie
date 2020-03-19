package main

import "fmt"

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	partitionIndex := start
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			// swap
			arr[i], arr[partitionIndex] = arr[partitionIndex], arr[i]
			partitionIndex++
		}
	}
	arr[partitionIndex], arr[end] = arr[end], arr[partitionIndex]
	return partitionIndex
}

func quickSort(arr []int, start, end int) {
	fmt.Println(arr)
	if start < end {
		partitionIndex := partition(arr, start, end)
		quickSort(arr, start, partitionIndex-1)
		quickSort(arr, partitionIndex+1, end)
	}
}

func main() {
	arr := []int{5, 4, 7, 8, 2, 4, 8, 9, 1}
	fmt.Println(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
