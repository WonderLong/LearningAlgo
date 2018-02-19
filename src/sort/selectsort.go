package main

import "fmt"

func SelectSort(arr []int) {
	leng := len(arr)
	for i := 0; i < leng; i++ {
		minIdx := i
		for j := i + 1; j < leng; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		swap(arr, i, minIdx)
	}
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	SelectSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d,", arr[i])
	}
	fmt.Println()
}
