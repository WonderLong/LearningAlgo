package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	QuickSort(arr)
	fmt.Println(arr)
}

/**
随机化优化快速排序提高效率

*/

func QuickSort(arr []int) {

	doSort(arr, 0, len(arr)-1)
}

func doSort(arr []int, l, r int) {

	if l >= r {
		return
	}

	rand.Seed(13)

	p := partition(arr, l, r)

	doSort(arr, l, p)
	doSort(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
	ra := rand.Intn(r - l)
	arr[l], arr[l+ra] = arr[l+ra], arr[l]
	v := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {

		if arr[i] < v {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j = j + 1
		}

	}

	arr[j], arr[l] = arr[l], arr[j]

	return j
}
