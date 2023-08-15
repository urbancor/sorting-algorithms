package main

import "fmt"

func insertion_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var j = 0
		for arr[j] < arr[i] && j < i {
			j++
		}
		for j <= i {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
}

func main() {
	var arr = []int{5, 4, 3, 2, 1}
	insertion_sort(arr)
	fmt.Println(arr)
}
