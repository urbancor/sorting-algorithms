package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ { // -i because the last i elements are already sorted
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	bubbleSort(arr)
	fmt.Println(arr)
}
