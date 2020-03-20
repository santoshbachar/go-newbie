package main

import "fmt"

func selectionSort(arr []int, size int) {
	for i := 0; i < size; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	nums := []int{3, 5, 1, 7, 2, 8}
	fmt.Println(nums)
	selectionSort(nums, len(nums)-1)
	fmt.Println(nums)
}
