package main

import "fmt"

func bubbleSort(arr []int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	nums := []int{8, 3, 2, 6, 3, 5, 1}
	fmt.Println(nums)
	bubbleSort(nums, len(nums)-1)
	fmt.Println(nums)
}
