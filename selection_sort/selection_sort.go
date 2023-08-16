package main

import "fmt"

func selection_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{1, 3, 2, 5, 4}
	selection_sort(arr)
	fmt.Println(arr)
}
